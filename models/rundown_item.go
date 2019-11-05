package models

type RundownItem struct {
	ID        int64  `json:id`
	Title     string `json:title`
	Subtitle  string `json:subtitle`
	Text      string `json:text`
	RundownId int64  `json:rundownId`
}
