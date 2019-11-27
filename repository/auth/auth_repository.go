package auth

import (
	"context"
	"database/sql"

	dto "../../dto"
	models "../../models"
	pRepo "../../repository"
)

func InitAuthRepository(Connection *sql.DB) pRepo.AuthRepository {
	return &AuthRepository{
		Connection: Connection,
	}
}

type AuthRepository struct {
	Connection *sql.DB
}

func (o *AuthRepository) Create(ctx context.Context, organizer *models.Organizer, user *models.User, account *models.Account) dto.Auth {
	organizerQuery := "Insert organizers SET name=?, description=?"
	userQuery := "Insert users SET name=?, organizer_id=?"
	accountQuery := "Insert accounts SET username=?, password=?, user_id=?"

	tx, _ := o.Connection.Begin()

	// ===========

	organizerStatement, err := tx.PrepareContext(ctx, organizerQuery)
	if err != nil {
		tx.Rollback()
	}

	organizerResponse, err := organizerStatement.ExecContext(ctx, organizer.Name, organizer.Description)
	defer organizerStatement.Close()
	if err != nil {
		tx.Rollback()
	}

	// ===========

	userStatement, err := tx.PrepareContext(ctx, userQuery)
	if err != nil {
		tx.Rollback()
	}

	organizerID, _ := organizerResponse.LastInsertId()
	organizer.ID = organizerID
	userResponse, err := userStatement.ExecContext(ctx, user.Name, organizerID)
	defer userStatement.Close()
	if err != nil {
		tx.Rollback()
	}

	// ===========

	accountStatement, err := tx.PrepareContext(ctx, accountQuery)
	if err != nil {
		tx.Rollback()
	}

	userID, _ := userResponse.LastInsertId()
	user.ID = userID
	accountResponse, err := accountStatement.ExecContext(ctx, account.Username, account.Password, userID)
	accountID, _ := accountResponse.LastInsertId()
	account.ID = accountID

	defer accountStatement.Close()
	if err != nil {
		tx.Rollback()
	}

	tx.Commit()

	authResponse := dto.Auth{
		Account:   account,
		User:      user,
		Organizer: organizer,
	}

	return authResponse
}
