package youtrack_parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/generat17/youtrack-parser/internal/config"
	"github.com/generat17/youtrack-parser/internal/models"
)

const (
	youTrackCurrentSprintApiUrl = "https://youtrack.wildberries.ru/api/agiles/<id>/sprints/current"
	sprintFields                = `fields=id,name,goal,start,finish,archived,isDefault,issues(<issue_fields>),unresolvedIssuesCount`

	youTrackGetIssueHistory = "https://youtrack.wildberries.ru/api/issues/<id>/activities"
	historyCategories       = `categories=<history_categories>`
	historyFields           = `fields=<history_fields>`

	typeIssueIndexInCustomFields            = 0
	stateIndexInCustomFields                = 1
	priorityIndexInCustomFields             = 2
	assigneeIndexInCustomFields             = 3
	taskAppearanceDateIndexInCustomFields   = 4
	deadlineIndexInCustomFields             = 5
	originalEstimationIndexInCustomFields   = 6
	completionPercentageIndexInCustomFields = 7
	tagsIndexInCustomFields                 = 8
	estimationIndexInCustomFields           = 9
	componentIndexInCustomFields            = 10
	spentTimeIndexInCustomFields            = 11

	configPath = "config.yaml"
)

var (
	conf config.Config
)

// GetSprint Возвращает данные по спринту
func GetSprint(youtrackApiToken string) (models.NormalizedSprint, error) {

	newConf, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("getting config: %v", err)
	}
	conf = newConf

	sprint, err := getSprint(conf.IssueFields, conf.AgileId, youtrackApiToken)
	if err != nil {
		return models.NormalizedSprint{}, fmt.Errorf("getting sprints: %v", err)
	}

	normalizedIssues := []models.NormalizedIssue{}

	// Мапа для фильтрации состояний
	states := make(map[string]struct{})
	for _, state := range conf.StatesWhitelist {
		states[state] = struct{}{}
	}

	for _, issue := range sprint.Issues {
		// проверка на черновик
		if conf.IsSkipDrafts && issue.IsDraft {
			continue
		}

		// получаем нормализованные кастомные поля
		normalizedCustomFields := convertCustomFields(issue.CustomFields)

		// проверка на наличие в state_whitelist
		_, ok := states[normalizedCustomFields.State.CustomValue.Name[0]]
		if !ok {
			continue
		}

		historyStateChanges := []models.NormalizedHistoryElementResponse{}
		historyCompletionPercentage := []models.NormalizedHistoryElementResponse{}
		historySpentTime := []models.NormalizedHistoryElementResponse{}

		// выполнение запроса к API YouTrack для получения истории изменений состояния
		history, err := getIssueHistory(conf.HistoryCategories, conf.HistoryFields, issue.IdReadable, youtrackApiToken)
		if err != nil {
			return models.NormalizedSprint{}, fmt.Errorf("getting history: %v", err)
		}

		for _, val := range history {
			current := parsingChangeHistory(val)

			if current.Field.Name == "% готово" {
				historyCompletionPercentage = append(historyCompletionPercentage, models.NormalizedHistoryElementResponse{
					Id:        val.Id,
					Added:     current.Added,
					Removed:   current.Removed,
					Author:    val.Author,
					Timestamp: val.Timestamp,
					Field:     val.Field,
				})
				continue
			}

			if current.Field.Name == "Spent time" {
				historySpentTime = append(historySpentTime, models.NormalizedHistoryElementResponse{
					Id:        val.Id,
					Added:     current.Added,
					Removed:   current.Removed,
					Author:    val.Author,
					Timestamp: val.Timestamp,
					Field:     val.Field,
				})
				continue
			}

			if current.Added.Type == "StateBundleElement" || current.Removed.Type == "StateBundleElement" {
				historyStateChanges = append(historyStateChanges, models.NormalizedHistoryElementResponse{
					Id:        val.Id,
					Added:     current.Added,
					Removed:   current.Removed,
					Author:    val.Author,
					Timestamp: val.Timestamp,
					Field:     val.Field,
				})
				continue
			}
		}

		normalizedIssues = append(normalizedIssues, models.NormalizedIssue{
			Id:                          issue.Id,
			Created:                     issue.Created,
			Description:                 issue.Description,
			IdReadable:                  issue.IdReadable,
			IsDraft:                     issue.IsDraft,
			NumberInProject:             issue.NumberInProject,
			Resolved:                    issue.Resolved,
			Summary:                     issue.Summary,
			Updated:                     issue.Updated,
			Votes:                       issue.Votes,
			WikifiedDescription:         issue.WikifiedDescription,
			CommentsCount:               issue.CommentsCount,
			Comments:                    issue.Comments,
			CustomFields:                normalizedCustomFields,
			Updater:                     issue.Updater,
			Links:                       issue.Links,
			HistoryStateChanges:         historyStateChanges,
			HistorySpentTime:            historySpentTime,
			HistoryCompletionPercentage: historyCompletionPercentage,
		})
	}

	normalizedSprintResponse := models.NormalizedSprint{
		Id:                    sprint.Id,
		Name:                  sprint.Name,
		Goal:                  sprint.Goal,
		Start:                 sprint.Start,
		Finish:                sprint.Finish,
		Archived:              sprint.Archived,
		IsDefault:             sprint.IsDefault,
		Issues:                normalizedIssues,
		UnresolvedIssuesCount: sprint.UnresolvedIssuesCount,
	}

	// записываем normalizedSprintResponse в json
	e, err := json.Marshal(&normalizedSprintResponse)
	if err != nil {
		return models.NormalizedSprint{}, fmt.Errorf("json marshal error: %v", err)
	}

	// запись json в файл
	err = writeToFile(e)
	if err != nil {
		return models.NormalizedSprint{}, fmt.Errorf("write to file error: %v", err)
	}

	return normalizedSprintResponse, nil
}

// GetListNormalNames возвращает fullname(YT): Нормальное имя, ссылка на тг
func GetListNormalNames() map[string][2]string {
	return conf.ListNormalNames
}

func writeToFile(jsonByteArr []byte) error {

	// создаем файл для вывода
	out, err := os.Create("sprint.json")
	if err != nil {
		return fmt.Errorf("write to file: %v", err)
	}
	// в конце программы, закрываем файл вывода
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	_, err = out.Write(jsonByteArr)
	if err != nil {
		return fmt.Errorf("write to file: %v", err)
	}

	return nil
}

// https://www.jetbrains.com/help/youtrack/devportal/resource-api-agiles-agileID-sprints.html
func getSprint(arrayFields []string, agileId, youtrackApiToken string) (models.SprintResponse, error) {
	path := strings.Replace(youTrackCurrentSprintApiUrl, "<id>", agileId, 1)
	fields := strings.Replace(sprintFields, "<issue_fields>", strings.Join(arrayFields, ","), 1)

	client := &http.Client{}
	req, err := http.NewRequest("GET", path+"?"+fields, nil)
	if err != nil {
		return models.SprintResponse{}, fmt.Errorf("sprint req: making new request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", youtrackApiToken))
	res, err := client.Do(req)
	if err != nil {
		return models.SprintResponse{}, fmt.Errorf("sprint req: getting sprint: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.SprintResponse{}, fmt.Errorf("sprint req: reading body: %v", err)
	}
	defer res.Body.Close()

	var sprints models.SprintResponse
	err = json.Unmarshal(body, &sprints)
	if err != nil {
		return models.SprintResponse{}, fmt.Errorf("sprint req: unmarshalling query body: %v", err)
	}

	return sprints, nil
}

// https://www.jetbrains.com/help/youtrack/devportal/api-usecase-issue-history.html
func getIssueHistory(arrayCategories []string, arrayFields []string, issueId, youtrackApiToken string) ([]models.HistoryElementResponse, error) {
	path := strings.Replace(youTrackGetIssueHistory, "<id>", issueId, 1)
	categories := strings.Replace(historyCategories, "<history_categories>", strings.Join(arrayCategories, ","), 1)
	fields := strings.Replace(historyFields, "<history_fields>", strings.Join(arrayFields, ","), 1)

	client := &http.Client{}
	req, err := http.NewRequest("GET", path+"?"+categories+"&"+fields, nil)
	if err != nil {
		return []models.HistoryElementResponse{}, fmt.Errorf("history req: making new request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", youtrackApiToken))
	res, err := client.Do(req)
	if err != nil {
		return []models.HistoryElementResponse{}, fmt.Errorf("history req: getting history: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []models.HistoryElementResponse{}, fmt.Errorf("history req: reading body: %v", err)
	}
	defer res.Body.Close()

	// подставить нужную модель
	var history []models.HistoryElementResponse
	err = json.Unmarshal(body, &history)
	if err != nil {
		return []models.HistoryElementResponse{}, fmt.Errorf("history req: unmarshalling query body: %v", err)
	}

	return history, nil
}

// парсит все кастомные поля issue в структуру models.NormalizedCustomFields
func convertCustomFields(customFieldsArr []models.CustomField) models.NormalizedCustomFields {
	return models.NormalizedCustomFields{
		TypeIssue: models.NormalizedCustomField{
			Id:          customFieldsArr[typeIssueIndexInCustomFields].Id,
			Name:        customFieldsArr[typeIssueIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[typeIssueIndexInCustomFields].Value),
		},
		State: models.NormalizedCustomField{
			Id:          customFieldsArr[stateIndexInCustomFields].Id,
			Name:        customFieldsArr[stateIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[stateIndexInCustomFields].Value),
		},
		Priority: models.NormalizedCustomField{
			Id:          customFieldsArr[priorityIndexInCustomFields].Id,
			Name:        customFieldsArr[priorityIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[priorityIndexInCustomFields].Value),
		},
		Assignee: models.NormalizedCustomField{
			Id:          customFieldsArr[assigneeIndexInCustomFields].Id,
			Name:        customFieldsArr[assigneeIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[assigneeIndexInCustomFields].Value),
		},
		TaskAppearanceDate: models.NormalizedCustomField{
			Id:          customFieldsArr[taskAppearanceDateIndexInCustomFields].Id,
			Name:        customFieldsArr[taskAppearanceDateIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[taskAppearanceDateIndexInCustomFields].Value),
		},
		Deadline: models.NormalizedCustomField{
			Id:          customFieldsArr[deadlineIndexInCustomFields].Id,
			Name:        customFieldsArr[deadlineIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[deadlineIndexInCustomFields].Value),
		},
		OriginalEstimation: models.NormalizedCustomField{
			Id:          customFieldsArr[originalEstimationIndexInCustomFields].Id,
			Name:        customFieldsArr[originalEstimationIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[originalEstimationIndexInCustomFields].Value),
		},
		CompletionPercentage: models.NormalizedCustomField{
			Id:          customFieldsArr[completionPercentageIndexInCustomFields].Id,
			Name:        customFieldsArr[completionPercentageIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[completionPercentageIndexInCustomFields].Value),
		},
		Tags: models.NormalizedCustomField{
			Id:          customFieldsArr[tagsIndexInCustomFields].Id,
			Name:        customFieldsArr[tagsIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[tagsIndexInCustomFields].Value),
		},
		Estimation: models.NormalizedCustomField{
			Id:          customFieldsArr[estimationIndexInCustomFields].Id,
			Name:        customFieldsArr[estimationIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[estimationIndexInCustomFields].Value),
		},
		Components: models.NormalizedCustomField{
			Id:          customFieldsArr[componentIndexInCustomFields].Id,
			Name:        customFieldsArr[componentIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[componentIndexInCustomFields].Value),
		},
		SpentTime: models.NormalizedCustomField{
			Id:          customFieldsArr[spentTimeIndexInCustomFields].Id,
			Name:        customFieldsArr[spentTimeIndexInCustomFields].Name,
			CustomValue: convertAnyToCustomField(customFieldsArr[spentTimeIndexInCustomFields].Value),
		},
	}
}

// Вспомогательная функция для функции convertCustomFields
// кастит любое кастомное значение к универсальному типу models.NormalizedCustomValue
func convertAnyToCustomField(value interface{}) models.NormalizedCustomValue {
	normalizedCustomValue := models.NormalizedCustomValue{
		Type:       "",
		Name:       []string{},
		FullName:   []string{},
		IsResolved: false,
		Timestamp:  0,
	}

	// пробую скастить значение в float64 (в случае если там лежит unix time)
	timestamp, isInt64 := value.(float64)
	if isInt64 {
		normalizedCustomValue.Timestamp = int64(timestamp)
	}

	// пробуем скастить в массив map
	arrayOfMaps, isArrayOfMaps := value.([]interface{})
	if isArrayOfMaps {
		for _, raw := range arrayOfMaps {
			rawMap, ok := raw.(map[string]interface{})
			if !ok {
				return models.NormalizedCustomValue{}
			}

			isResolved, ok := rawMap["isResolved"]
			if ok {
				normalizedCustomValue.IsResolved = isResolved.(bool)
			}

			typeValue, ok := rawMap["$type"]
			if ok {
				normalizedCustomValue.Type = typeValue.(string)
			}

			// name parsing
			nameAny, ok := rawMap["name"]
			if ok {
				name := nameAny.(string)
				normalizedCustomValue.Name = append(normalizedCustomValue.Name, name)
			}

			// fullName parsing
			fullNameAny, ok := rawMap["fullName"]
			if ok {
				fullName := fullNameAny.(string)
				normalizedCustomValue.FullName = append(normalizedCustomValue.FullName, fullName)
			}
		}
	}

	// пробуем скастить в мапу
	singleMap, isSingleMap := value.(map[string]interface{})
	if isSingleMap {
		isResolved, ok := singleMap["isResolved"]
		if ok {
			normalizedCustomValue.IsResolved = isResolved.(bool)
		}

		typeValue, ok := singleMap["$type"]
		if ok {
			normalizedCustomValue.Type = typeValue.(string)
		}

		// name parsing
		nameAny, ok := singleMap["name"]
		if ok {
			name := nameAny.(string)
			normalizedCustomValue.Name = append(normalizedCustomValue.Name, name)
		}

		// fullName parsing
		fullNameAny, ok := singleMap["fullName"]
		if ok {
			fullName := fullNameAny.(string)
			normalizedCustomValue.FullName = append(normalizedCustomValue.FullName, fullName)
		}
	}

	return normalizedCustomValue
}

// парсит элемент истории
func parsingChangeHistory(historyElement models.HistoryElementResponse) models.NormalizedHistoryElementResponse {
	return models.NormalizedHistoryElementResponse{
		Id:        historyElement.Id,
		Added:     convertAnyToCustomField(historyElement.Added),
		Removed:   convertAnyToCustomField(historyElement.Removed),
		Author:    historyElement.Author,
		Timestamp: historyElement.Timestamp,
		Field:     historyElement.Field,
	}
}
