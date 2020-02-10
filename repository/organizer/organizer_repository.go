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
			&data.DisplayName,
			&data.Description,
			&data.LocationName,
			&data.LocationLat,
			&data.LocationLng,
			&data.LocationAddress,
			&data.City,
			&data.Nation,
		)

		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *OrganizerRepository) GetByID(ctx context.Context, id int64) (*models.Organizer, error) {
	query := "Select * From organizers where id=?"

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
	query := "Select * From organizers where name LIKE '" + "%" + name + "%" + "' or display_name LIKE '" + "%" + name + "%" + "'"

	return m.fetch(ctx, query)
}

func (m *OrganizerRepository) GetByCity(ctx context.Context, city string) ([]*models.Organizer, error) {
	query := "Select * From organizers where city LIKE '" + "%" + city + "%" + "'"

	return m.fetch(ctx, query)
}

func (m *OrganizerRepository) GetByProvince(ctx context.Context, province string) ([]*models.Organizer, error) {
	query := "Select * From organizers where province LIKE '" + "%" + province + "%" + "'"

	return m.fetch(ctx, query)
}
