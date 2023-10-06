package models

// https://www.jetbrains.com/help/youtrack/devportal/api-entity-IssueComment.html

type IssueComment struct {
	Id          string `json:"id"`
	Text        string `json:"text"`
	TextPreview string `json:"textPreview"`
	Created     int64  `json:"created"`
	Updated     int64  `json:"updated"`
	Author      User   `json:"author"`
	Deleted     bool   `json:"deleted"`

	// Unused fields

	// Attachments []IssueAttachment `json:"attachments"`
	// Issue       Issue             `json:"issue"`
	// Visibility interface{}		`json:"visibility"`
}
