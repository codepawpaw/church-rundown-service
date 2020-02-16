package models

import "time"

type ServiceSchedule struct {
	ID          int64     `json:id`
	Name        string    `json:name`
	Text        string    `json:text`
	Date        time.Time `json:date`
	OrganizerId int64     `json:"organizerID"`
}
