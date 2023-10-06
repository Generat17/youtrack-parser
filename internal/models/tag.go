package models

// https://www.jetbrains.com/help/youtrack/devportal/api-entity-Tag.html

type Tag struct {
	Id             string  `json:"id"`
	Issue          []Issue `json:"issue"`
	UntagOnResolve bool    `json:"untagOnResolve"`
	Owner          User    `json:"owner"`
	Name           string  `json:"name"`

	// Unused fields

	// Color                 interface{} `json:"color"`
	// VisibleFor            interface{} `json:"visibleFor"`
	// UpdateableBy          interface{} `json:"updateableBy"`
	// ReadSharingSettings   interface{} `json:"readSharingSettings"`
	// TagSharingSettings    interface{} `json:"tagSharingSettings"`
	// UpdateSharingSettings interface{} `json:"updateSharingSettings"`
}
