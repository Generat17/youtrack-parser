package models

// https://www.jetbrains.com/help/youtrack/devportal/resource-api-agiles-agileID-sprints.html

type SprintResponse struct {
	Id                    string  `json:"id"`
	Name                  string  `json:"name"`
	Goal                  string  `json:"goal"`
	Start                 int64   `json:"start"`
	Finish                int64   `json:"finish"`
	Archived              bool    `json:"archived"`
	IsDefault             bool    `json:"isDefault"`
	Issues                []Issue `json:"issues"`
	UnresolvedIssuesCount int     `json:"unresolvedIssuesCount"`

	// Unused fields
	// Agile          interface{} `json:"agile"`
	// PreviousSprint interface{} `json:"previousSprint"`
}

type NormalizedSprint struct {
	Id                    string            `json:"id"`
	Name                  string            `json:"name"`
	Goal                  string            `json:"goal"`
	Start                 int64             `json:"start"`
	Finish                int64             `json:"finish"`
	Archived              bool              `json:"archived"`
	IsDefault             bool              `json:"isDefault"`
	Issues                []NormalizedIssue `json:"issues"`
	UnresolvedIssuesCount int               `json:"unresolvedIssuesCount"`

	// Unused fields
	// Agile          interface{} `json:"agile"`
	// PreviousSprint interface{} `json:"previousSprint"`
}
