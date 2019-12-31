package organizer

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitOrganizerRepository(Connection *sql.DB) pRepo.OrganizerRepository {
	return &OrganizerRepository{
		Connection: Connection,
	}
}

type OrganizerRepository struct {
	Connection *sql.DB
}

func (o *OrganizerRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Organizer, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Organizer, 0)
	for rows.Next() {
		data := new(models.Organizer)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Description,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (o *OrganizerRepository) GetAll(ctx context.Context, num int64, id string, name string) ([]*models.Organizer, error) {
	query := "Select id, name, description From organizers"
	counter := 0

	if id != "" {
		query = query + " where id=?"
		counter += 1
	}

	if name != "" {
		if counter > 0 {
			query = query + " AND name LIKE '" + "%" + name + "%" + "' limit ?"
			return o.fetch(ctx, query, id, num)
		} else {
			query = query + " where name LIKE '" + "%" + name + "%" + "' limit ?"
			return o.fetch(ctx, query, num)
		}
	}

	query = query + " limit ?"
	if counter > 0 {
		return o.fetch(ctx, query, id, num)
	}

	return o.fetch(ctx, query, num)
}

func (o *OrganizerRepository) Create(ctx context.Context, p *models.Organizer) (int64, error) {
	query := "Insert organizers SET name=?, description=?"

	stmt, err := o.Connection.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Name, p.Description)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *OrganizerRepository) GetByID(ctx context.Context, id int64) (*models.Organizer, error) {
	query := "Select id, name, description From organizers where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Organizer{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *OrganizerRepository) GetByName(ctx context.Context, name string) ([]*models.Organizer, error) {
	query := "Select id, name, description From organizers where name like ?"

	return m.fetch(ctx, query, name)
}

func (m *OrganizerRepository) Update(ctx context.Context, p *models.Organizer) (*models.Organizer, error) {
	query := "Update organizers set name=?, description=? where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.Description,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *OrganizerRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From organizers Where id=?"

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
