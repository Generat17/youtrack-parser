package config

var (
	AgileId         = "83-1793"
	IsSkipDrafts    = true
	StatesWhitelist = []string{
		"Backlog",
		"Canceled",
		"Deployed",
		"Done",
		"In Deployment",
		"In Progress",
		"In Review",
		"Open",
		"Testing",
	}
	IssueFields = []string{
		"id",
		"created",
		"description",
		"idReadable",
		"isDraft",
		"numberInProject",
		"resolved",
		"summary",
		"updated",
		"votes",
		"wikifiedDescription",
		"commentsCount",
		"comments(id,text,textPreview,created,updated,deleted,author(id,fullName,email))",
		"customFields(id,name,value($type,name,fullName,isResolved))",
		"updater(id,fullName,email)",
		"tags(id,name)",
		"reporter(id,fullName,email)",
		"draftOwner(id,fullName,email)",
		"links(id,direction,linkType(id,name,localizedName,sourceToTarget,localizedSourceToTarget,targetToSource,localizedTargetToSource,directed,aggregation,readOnly),issues(id,idReadable),trimmedIssues(id,idReadable))",
	}
	HistoryCategories = []string{
		"CustomFieldCategory",
	}
	HistoryFields = []string{
		"id",
		"timestamp",
		"author(name,login)",
		"added(id,name,value($type,name,fullName,isResolved))",
		"removed(id,name,value($type,name,fullName,isResolved))",
		"target",
		"category",
		"field(id,presentation,name)",
	}
	ListNormalNames = map[string][2]string{
		"Anna Peremitina":     {"Перемитина Анна", "https://t.me/sandra_kas_sandra"},
		"David Utyuganov":     {"Утюганов Давид", "https://t.me/singeroux"},
		"Denis Baranov":       {"Баранов Денис", "https://t.me/denis319199"},
		"Dmitry Deev":         {"Деев Дмитрий", "https://t.me/dinvictus"},
		"Grant Simonyan":      {"Симонян Грант", "https://t.me/sgurman"},
		"Ivan Chayka":         {"Чайка Иван", "https://t.me/iv0xff"},
		"Ivan_Gurianov":       {"Гурьянов Иван", "https://t.me/Jackalivan"},
		"Kucherov Konstantin": {"Кучеров Константин", "https://t.me/Expresso_const"},
		"Nikita Bulchuk":      {"Бульчук Никита", "https://t.me/nikrossin"},
		"Parviz Dzhamilov":    {"Джамилов Парвиз", "https://t.me/parvizjamilov"},
		"Popov Nikita":        {"Попов Никита", "https://t.me/n_priest"},
		"Ramil Nafeev":        {"Нафеев Рамиль", "https://t.me/Pramill"},
		"Shamil Mukhetdinov":  {"Мухетдинов Шамиль", "https://t.me/ShamilC137"},
		"baryshnikov.n9":      {"Барышников Никита", "https://t.me/shushard"},
		"galushkin.d4":        {"Галушкин Дмитрий", "https://t.me/dimyasha"},
		"gubeydullin.i":       {"Губейдуллин Ильнур", "https://t.me/lilkaide"},
		"istamov.valeriy":     {"Истамов Валерий", "https://t.me/istamov_valery"},
		"kabanov.denis4":      {"Кабанов Денис", "https://t.me/DionisiusOfAcadem"},
		"korolev.oleg7":       {"Королев Олег", "https://t.me/kor0lll"},
		"kryachko.v2":         {"Крячко Владимир", "https://t.me/destinyxus"},
		"nepryahin.mihail":    {"Непряхин Михаил", "https://t.me/neprja"},
		"romanenko.denis4":    {"Романенко Денис", "https://t.me/TheStraitor"},
		"ryakin.pavel":        {"Рякин Павел", "https://t.me/pasharyakin"},
		"suluyanov.egor":      {"Сулуянов Егор", "https://t.me/suluyanove"},
		"Алиев Тимур":         {"Алиев Тимур", "https://t.me/Aliev_Timur_M"},
		"Владислав Косогоров": {"Косогоров Владислав", "https://t.me/ketsuwotaberu"},
		"Вячеслав Киндеев":    {"Киндеев Вячеслав", "https://t.me/xqqzw"},
		"Николаев Владислав Андреевич": {"Николаев Владислав", "https://t.me/dlavkonievla"},
		"Павел Николаевский":           {"Николаевский Павел", "https://t.me/Utrian"},
		"Павленко Владислав Вадимович": {"Павленко Владислав", "https://t.me/pavlenkowild"},
		"Устелемов Максим Алексеевич":  {"Устелемов Максим", "https://t.me/maksustoff"},
		"Эльгин Сергей Михайлович":     {"Эльгин Сергей", "https://t.me/AiDiod"},
	}
)

type Config struct {
	AgileId           string               `yaml:"agile_id"`
	IsSkipDrafts      bool                 `yaml:"skip_drafts"`
	StatesWhitelist   []string             `yaml:"state_whitelist"`
	IssueFields       []string             `yaml:"issue_fields"`
	HistoryCategories []string             `yaml:"history_categories"`
	HistoryFields     []string             `yaml:"history_fields"`
	ListNormalNames   map[string][2]string `yaml:"normal_names"`
}

func NewConfig() (Config, error) {
	return readConfigFile()
}

func readConfigFile() (Config, error) {
	config := Config{
		AgileId:           AgileId,
		IsSkipDrafts:      IsSkipDrafts,
		StatesWhitelist:   StatesWhitelist,
		IssueFields:       IssueFields,
		HistoryCategories: HistoryCategories,
		HistoryFields:     HistoryFields,
		ListNormalNames:   ListNormalNames,
	}

	return config, nil
}
