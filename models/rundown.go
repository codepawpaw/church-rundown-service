package models

import "time"

type Rundown struct {
	ID          int64     `json:id`
	Title       string    `json:title`
	Subtitle    string    `json:subtitle`
	ShowTime    time.Time `json:"showTime"`
	EndTime     time.Time `json:"endTime"`
	OrganizerId int64     `json:organizerId`
}
