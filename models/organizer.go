package models

type Organizer struct {
	ID          int64  `json:id`
	Name        string `json:name`
	Description string `json:description`
}