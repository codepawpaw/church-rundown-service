package models

import "time"

type DeviceInventory struct {
	ID           int64     `json:id`
	Name         string    `json:name`
	PurchaseDate time.Time `json:purchaseDate`
	Total        int64     `json:"total"`
	OrganizerId  int64     `json:"organizerID"`
}
