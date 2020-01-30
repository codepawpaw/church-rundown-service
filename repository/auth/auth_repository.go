package auth

import (
	"context"
	"database/sql"
	"fmt"

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

func (o *AuthRepository) Create(ctx context.Context, organizer *models.Organizer, user *models.User, account *models.Account) (dto.Auth, error) {
	organizerQuery := "Insert organizers SET name=?, description=?, location_name=?, location_lat=?, location_lng=?, location_address=?"
	userQuery := "Insert users SET name=?, organizer_id=?"
	accountQuery := "Insert accounts SET username=?, password=?, user_id=?"

	tx, _ := o.Connection.Begin()

	emptyAuthResponse := dto.Auth{}

	// ===========

	organizerStatement, err := tx.PrepareContext(ctx, organizerQuery)
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	organizerResponse, err := organizerStatement.ExecContext(ctx, organizer.Name, organizer.Description, organizer.LocationName, organizer.LocationLat, organizer.LocationLng, organizer.LocationAddress)
	defer organizerStatement.Close()
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	// ===========

	userStatement, err := tx.PrepareContext(ctx, userQuery)
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	organizerID, _ := organizerResponse.LastInsertId()
	organizer.ID = organizerID
	userResponse, err := userStatement.ExecContext(ctx, user.Name, organizerID)
	defer userStatement.Close()
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	// ===========

	accountStatement, err := tx.PrepareContext(ctx, accountQuery)
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	userID, _ := userResponse.LastInsertId()
	user.ID = userID
	accountResponse, err := accountStatement.ExecContext(ctx, account.Username, account.Password, userID)
	defer accountStatement.Close()

	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	accountID, _ := accountResponse.LastInsertId()
	account.ID = accountID

	tx.Commit()

	authResponse := dto.Auth{
		Account:   account,
		User:      user,
		Organizer: organizer,
	}

	return authResponse, err
}

func (o *AuthRepository) Update(ctx context.Context, organizer *models.Organizer, user *models.User, account *models.Account) (dto.Auth, error) {
	organizerQuery := "Update organizers SET name=?, description=? WHERE id=?"
	userQuery := "Update users SET name=? Where id=?"
	accountQuery := "Update accounts SET username=?, password=? Where id=?"

	tx, _ := o.Connection.Begin()

	emptyAuthResponse := dto.Auth{}

	// ===========

	organizerStatement, err := tx.PrepareContext(ctx, organizerQuery)
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	organizerResponse, err := organizerStatement.ExecContext(ctx, organizer.Name, organizer.Description, organizer.ID)
	defer organizerStatement.Close()
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	// ===========

	userStatement, err := tx.PrepareContext(ctx, userQuery)
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	fmt.Println(organizerResponse)

	userResponse, err := userStatement.ExecContext(ctx, user.Name, user.ID)
	defer userStatement.Close()
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	// ===========

	accountStatement, err := tx.PrepareContext(ctx, accountQuery)
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	fmt.Println(userResponse)

	accountResponse, err := accountStatement.ExecContext(ctx, account.Username, account.Password, account.ID)
	defer accountStatement.Close()

	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	fmt.Println(accountResponse)

	tx.Commit()

	authResponse := dto.Auth{
		Account:   account,
		User:      user,
		Organizer: organizer,
	}

	return authResponse, err
}
