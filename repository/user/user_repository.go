package user

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitUserRepository(Connection *sql.DB) pRepo.UserRepository {
	return &UserRepository{
		Connection: Connection,
	}
}

type UserRepository struct {
	Connection *sql.DB
}

func (o *UserRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.User, 0)
	for rows.Next() {
		data := new(models.User)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.OrganizerId,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (o *UserRepository) GetAll(ctx context.Context, num int64) ([]*models.User, error) {
	query := "Select id, name, organizer_id From users limit ?"

	return o.fetch(ctx, query, num)
}

func (o *UserRepository) Create(ctx context.Context, p *models.User) (int64, error) {
	query := "Insert users SET name=?, organizer_id=?"

	stmt, err := o.Connection.PrepareContext(ctx, query)

	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Name, p.OrganizerId)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *UserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := "Select id, name, organizer_id From users where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.User{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *UserRepository) Update(ctx context.Context, p *models.User) (*models.User, error) {
	query := "Update users set name=? where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *UserRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From users Where id=?"

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
