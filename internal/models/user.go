package models

// https://www.jetbrains.com/help/youtrack/devportal/api-entity-User.html

type User struct {
	Id        string `json:"id"`
	Login     string `json:"login"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	RingId    string `json:"ringId"`
	Quest     bool   `json:"guest"`
	Online    bool   `json:"online"`
	Banned    bool   `json:"banned"`
	Tags      []Tag  `json:"tags"`
	AvatarUrl string `json:"avatarUrl"`
	Name      string `json:"name"` // атрибут не прописан в API, но historyElement -> author возрващает name, вместо fullName

	// Unused fields

	// SavedQueries interface{} `json:"savedQueries"` // заглушка
	// Profiles     interface{} `json:"profiles"` // заглушка
}
