package models

// https://www.jetbrains.com/help/youtrack/devportal/api-entity-FilterField.html

type FilterField struct {
	Id            string `json:"id"`
	Presentations string `json:"presentations"`
	Name          string `json:"name"`
}
