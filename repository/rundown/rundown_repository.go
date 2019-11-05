package rundown

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitRundownRepository(Connection *sql.DB) pRepo.RundownRepository {
	return &RundownRepository{
		Connection: Connection,
	}
}

type RundownRepository struct {
	Connection *sql.DB
}

func (o *RundownRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Rundown, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Rundown, 0)
	for rows.Next() {
		data := new(models.Rundown)

		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.Subtitle,
			&data.ShowTime,
			&data.EndTime,
			&data.OrganizerId,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (o *RundownRepository) GetAll(ctx context.Context, num int64) ([]*models.Rundown, error) {
	query := "Select id, title, subtitle, show_time, end_time, organizer_id From rundowns limit ?"

	return o.fetch(ctx, query, num)
}

func (o *RundownRepository) Create(ctx context.Context, p *models.Rundown) (int64, error) {
	query := "Insert rundowns SET title=?, subtitle=?, show_time=?, end_time=?, organizer_id=?"

	stmt, err := o.Connection.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Title, p.Subtitle, p.ShowTime, p.EndTime, p.OrganizerId)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *RundownRepository) GetByID(ctx context.Context, id int64) (*models.Rundown, error) {
	query := "Select * From rundowns where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Rundown{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *RundownRepository) GetByOrganizerId(ctx context.Context, organizerId int64, startDate string, endDate string) ([]*models.Rundown, error) {
	query := "Select * From rundowns where organizer_id=?"

	if startDate != "" && endDate != "" {
		query = query + " AND show_time >= ? AND end_time <= ?"
		return m.fetch(ctx, query, organizerId, startDate, endDate)
	}

	return m.fetch(ctx, query, organizerId)
}

func (m *RundownRepository) GetByOrganizerIdAndDate(ctx context.Context, organizerId int64, startDate string, endDate string) ([]*models.Rundown, error) {
	query := "Select * From rundowns where organizer_id=? and show_time >= ? and end_time <= ?"

	return m.fetch(ctx, query, organizerId, startDate, endDate)
}

func (m *RundownRepository) GetByOrganizerAndId(ctx context.Context, id int64, organizerId int64) (*models.Rundown, error) {
	query := "Select * from rundowns where id=? AND organizer_id=?"
	rows, err := m.fetch(ctx, query, id, organizerId)

	if err != nil {
		return nil, err
	}

	payload := &models.Rundown{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *RundownRepository) Update(ctx context.Context, p *models.Rundown) (*models.Rundown, error) {
	query := "Update rundowns set title=?, subtitle=?, show_time=?, end_time=? where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Title,
		p.Subtitle,
		p.ShowTime,
		p.EndTime,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *RundownRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From rundowns Where id=?"

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
