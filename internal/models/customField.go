package models

type CustomField struct {
	Id    string      `json:"id"`
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type NormalizedCustomFields struct {
	TypeIssue            NormalizedCustomField `json:"typeIssue"`
	State                NormalizedCustomField `json:"state"`
	Priority             NormalizedCustomField `json:"priority"`
	Assignee             NormalizedCustomField `json:"assignee"`
	TaskAppearanceDate   NormalizedCustomField `json:"taskAppearanceDate"`
	Deadline             NormalizedCustomField `json:"deadline"`
	OriginalEstimation   NormalizedCustomField `json:"originalEstimation"`
	CompletionPercentage NormalizedCustomField `json:"completionPercentage"`
	Tags                 NormalizedCustomField `json:"tags"`
	Estimation           NormalizedCustomField `json:"estimation"`
	Components           NormalizedCustomField `json:"components"`
	SpentTime            NormalizedCustomField `json:"spentTime"`
}

type NormalizedCustomField struct {
	Id          string                `json:"id"`
	Name        string                `json:"name"`
	CustomValue NormalizedCustomValue `json:"value"`
}

type NormalizedCustomValue struct {
	Type       string   `json:"type"`
	Name       []string `json:"name"`
	FullName   []string `json:"fullName"`
	IsResolved bool     `json:"isResolved"`
	Timestamp  int64    `json:"timestamp"`
}
