package models

type IssueLinkType struct {
	Id                      string `json:"id"`
	Name                    string `json:"name"`
	LocalizedName           string `json:"localizedName"`
	SourceToTarget          string `json:"sourceToTarget"`
	LocalizedSourceToTarget string `json:"localizedSourceToTarget"`
	TargetToSource          string `json:"targetToSource"`
	LocalizedTargetToSource string `json:"localizedTargetToSource"`
	Directed                bool   `json:"directed"`
	Aggregation             bool   `json:"aggregation"`
	ReadOnly                bool   `json:"readOnly"`
}
