package models

type Concregation struct {
	ID          int64  `json:id`
	Name        string `json:name`
	Age         string `json:age`
	Address     string `json:"address"`
	OrganizerId int64  `json:"organizerID"`
}
