package models

type Auth struct {
	Organizer Organizer `json:organizer`
	User      User      `json:user`
	Account   Account   `json:account`
}
