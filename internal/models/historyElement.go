package models

type HistoryElementResponse struct {
	Id        string      `json:"id"`
	Added     interface{} `json:"added"`
	Removed   interface{} `json:"removed"`
	Author    User        `json:"author"`
	Timestamp int64       `json:"timestamp"`
	Field     FilterField `json:"field"`

	// Unused fields

	//Category     interface{} `json:"category"`
	//TargetMember interface{} `json:"targetMember"`
	//Target  int         `json:"target"`
}

type NormalizedHistoryElementResponse struct {
	Id        string                `json:"id"`
	Added     NormalizedCustomValue `json:"added"`
	Removed   NormalizedCustomValue `json:"removed"`
	Author    User                  `json:"author"`
	Timestamp int64                 `json:"timestamp"`
	Field     FilterField           `json:"field"`

	// Unused fields

	//Category     interface{} `json:"category"`
	//TargetMember interface{} `json:"targetMember"`
	//Target  int         `json:"target"`
}
