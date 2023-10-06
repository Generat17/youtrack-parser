package models

// https://www.jetbrains.com/help/youtrack/devportal/api-entity-IssueAttachment.html

type IssueAttachment struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Author        User   `json:"author"`
	Created       int64  `json:"created"`
	Updated       int64  `json:"updated"`
	Size          int    `json:"size"`
	Extension     string `json:"extension"`
	Charset       string `json:"charset"`
	MimeType      string `json:"mimeType"`
	MetaData      string `json:"metaData"`
	Draft         bool   `json:"draft"`
	Removed       bool   `json:"removed"`
	Base64Content string `json:"base64Content"`
	Url           string `json:"url"`
	ThumbnailURL  string `json:"thumbnailURL"`

	// Unused fields

	// Visibility    interface{} `json:"visibility"`
	// Issue         interface{} `json:"issue"`
	// Comment       interface{} `json:"comment"`
}
