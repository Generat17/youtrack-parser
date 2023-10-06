package models

type IssueLink struct {
	Id            string        `json:"id"`
	Direction     string        `json:"direction"`
	LinkType      IssueLinkType `json:"linkType"`
	Issues        []Issue       `json:"issues"`
	TrimmedIssues []Issue       `json:"trimmedIssues"`
}
