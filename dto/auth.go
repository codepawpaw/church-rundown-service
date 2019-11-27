package dto

import (
	models "../models"
)

type Auth struct {
	Account   *models.Account   `json:"account"`
	User      *models.User      `json:"user"`
	Organizer *models.Organizer `json:"organizer"`
	Token     string            `json:"token"`
}
