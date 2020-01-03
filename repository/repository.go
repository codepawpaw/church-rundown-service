package repsitory

import (
	"context"

	dto "../dto"
	models "../models"
)

type OrganizerRepository interface {
	GetByID(ctx context.Context, id int64) (*models.Organizer, error)
	GetByName(ctx context.Context, name string) ([]*models.Organizer, error)
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
}

type AccountRepository interface {
	GetByID(ctx context.Context, id int64) (*models.Account, error)
	GetByUsernameAndPassword(ctx context.Context, username string, password string) (*models.Account, error)
	Update(ctx context.Context, p *models.Account) (*models.Account, error)
}

type RundownRepository interface {
	Create(ctx context.Context, p *models.Rundown) (*models.Rundown, error)
	GetByID(ctx context.Context, id int64) (*models.Rundown, error)
	GetByOrganizerId(ctx context.Context, organizerId int64, startDate string, endDate string) ([]*models.Rundown, error)
	GetByOrganizerIdAndDate(ctx context.Context, organizerId int64, startDate string, endDate string) ([]*models.Rundown, error)
	Update(ctx context.Context, p *models.Rundown) (*models.Rundown, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

type RundownItemRepository interface {
	Create(ctx context.Context, p *models.RundownItem) (*models.RundownItem, error)
	GetByRundownId(ctx context.Context, rundownid int64) ([]*models.RundownItem, error)
	Update(ctx context.Context, p *models.RundownItem) (*models.RundownItem, error)
	Delete(ctx context.Context, id int64) (bool, error)
	DeleteByRundownId(ctx context.Context, rundownId int64) (bool, error)
}

type AuthRepository interface {
	Create(ctx context.Context, organizer *models.Organizer, user *models.User, account *models.Account) (dto.Auth, error)
	Update(ctx context.Context, organizer *models.Organizer, user *models.User, account *models.Account) (dto.Auth, error)
}
