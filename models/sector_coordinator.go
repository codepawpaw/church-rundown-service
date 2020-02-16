package models

type SectorCoordinator struct {
	ID             int64  `json:id`
	Name           string `json:name`
	ConcregationId int64  `json:"concregationID"`
	OrganizerId    int64  `json:"organizerID"`
}
