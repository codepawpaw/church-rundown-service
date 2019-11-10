package account

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitAccountRepository(Connection *sql.DB) pRepo.AccountRepository {
	return &AccountRepository{
		Connection: Connection,
	}
}

type AccountRepository struct {
	Connection *sql.DB
}

func (o *AccountRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Account, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Account, 0)
	for rows.Next() {
		data := new(models.Account)

		err := rows.Scan(
			&data.ID,
			&data.Username,
			&data.Password,
			&data.UserId,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (o *AccountRepository) GetAll(ctx context.Context, num int64) ([]*models.Account, error) {
	query := "Select id, username, password, user_id From accounts limit ?"

	return o.fetch(ctx, query, num)
}

func (o *AccountRepository) Create(ctx context.Context, p *models.Account) (int64, error) {
	query := "Insert accounts SET username=?, password=?, user_id=?"

	stmt, err := o.Connection.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Username, p.Password, p.UserId)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *AccountRepository) GetByID(ctx context.Context, id int64) (*models.Account, error) {
	query := "Select id, username, password, user_id From accounts where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Account{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *AccountRepository) GetByUsernameAndPassword(ctx context.Context, username string, password string) (*models.Account, error) {
	query := "Select id, username, password, user_id From accounts where username=? and password=?"

	rows, err := m.fetch(ctx, query, username, password)
	if err != nil {
		return nil, err
	}

	payload := &models.Account{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *AccountRepository) GetByUsername(ctx context.Context, username string) (*models.Account, error) {
	query := "Select id, username, password, user_id From accounts where username=?"

	rows, err := m.fetch(ctx, query, username)
	if err != nil {
		return nil, err
	}

	payload := &models.Account{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *AccountRepository) Update(ctx context.Context, p *models.Account) (*models.Account, error) {
	query := "Update accounts set username=?, password=? where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Username,
		p.Password,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *AccountRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From accounts Where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
