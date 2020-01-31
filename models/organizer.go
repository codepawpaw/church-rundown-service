package models

type Organizer struct {
	ID              int64  `json:id`
	Name            string `json:name`
	DisplayName     string `json:"displayName"`
	Description     string `json:description`
	LocationName    string `json:"locationName"`
	LocationLat     string `json:"locationLat"`
	LocationLng     string `json:"locationLng"`
	LocationAddress string `json:"locationAddress"`
}
