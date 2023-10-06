package models

// https://www.jetbrains.com/help/youtrack/devportal/api-entity-Issue.html

type Issue struct {
	Id                  string         `json:"id"`
	Created             int            `json:"created"`
	Description         string         `json:"description"`
	IdReadable          string         `json:"idReadable"`
	IsDraft             bool           `json:"isDraft"`
	NumberInProject     int            `json:"numberInProject"`
	Resolved            int            `json:"resolved"`
	Summary             string         `json:"summary"`
	Updated             int64          `json:"updated"`
	Votes               int            `json:"votes"`
	WikifiedDescription string         `json:"wikifiedDescription"`
	CommentsCount       int            `json:"commentsCount"`
	Comments            []IssueComment `json:"comments"`
	CustomFields        []CustomField  `json:"customFields"`
	Updater             User           `json:"updater"`
	Tags                []Tag          `json:"tags"`
	Reporter            User           `json:"reporter"`
	DraftOwner          User           `json:"draftOwner"`
	Links               []IssueLink    `json:"links"`

	// Unused fields

	// Attachments         []IssueAttachment `json:"attachments"`
	// ExternalIssue       interface{}       `json:"externalIssue"`
	// Parent              interface{}       `json:"parent"`
	// Project             interface{}       `json:"project"`
	// Subtasks            interface{}       `json:"subtasks"`
	// Visibility          interface{}       `json:"visibility"`
	// Voters              interface{}       `json:"voters"`
	// Watchers            interface{}       `json:"watchers"`
}

type NormalizedIssue struct {
	Id                          string                             `json:"id"`
	Created                     int                                `json:"created"`
	Description                 string                             `json:"description"`
	IdReadable                  string                             `json:"idReadable"`
	IsDraft                     bool                               `json:"isDraft"`
	NumberInProject             int                                `json:"numberInProject"`
	Resolved                    int                                `json:"resolved"`
	Summary                     string                             `json:"summary"`
	Updated                     int64                              `json:"updated"`
	Votes                       int                                `json:"votes"`
	WikifiedDescription         string                             `json:"wikifiedDescription"`
	CommentsCount               int                                `json:"commentsCount"`
	Comments                    []IssueComment                     `json:"comments"`
	CustomFields                NormalizedCustomFields             `json:"customFields"`
	Updater                     User                               `json:"updater"`
	Tags                        []Tag                              `json:"tags"`
	Reporter                    User                               `json:"reporter"`
	DraftOwner                  User                               `json:"draftOwner"`
	Links                       []IssueLink                        `json:"links"`
	HistoryStateChanges         []NormalizedHistoryElementResponse `json:"historyStateChanges"`
	HistoryCompletionPercentage []NormalizedHistoryElementResponse `json:"historyCompletionPercentage"`
	HistorySpentTime            []NormalizedHistoryElementResponse `json:"historySpentTime"`

	// Unused fields

	// Attachments         []IssueAttachment `json:"attachments"`
	// ExternalIssue       interface{}       `json:"externalIssue"`
	// Parent              interface{}       `json:"parent"`
	// Project             interface{}       `json:"project"`
	// Subtasks            interface{}       `json:"subtasks"`
	// Visibility          interface{}       `json:"visibility"`
	// Voters              interface{}       `json:"voters"`
	// Watchers            interface{}       `json:"watchers"`
}
