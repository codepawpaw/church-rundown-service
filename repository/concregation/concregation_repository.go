package concregation

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitConcregationRepository(Connection *sql.DB) pRepo.ConcregationRepository {
	return &ConcregationRepository{
		Connection: Connection,
	}
}

type ConcregationRepository struct {
	Connection *sql.DB
}

func (o *ConcregationRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Concregation, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Concregation, 0)
	for rows.Next() {
		data := new(models.Concregation)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Age,
			&data.Address,
			&data.OrganizerId,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (o *ConcregationRepository) Create(ctx context.Context, concregation *models.Concregation) (*models.Concregation, error) {
	query := "Insert concregation SET name=?, age=?, address=?, organizer_id=?"

	stmt, err := o.Connection.PrepareContext(ctx, query)
	if err != nil {
		return &models.Concregation{}, err
	}

	concregationResponse, err := stmt.ExecContext(ctx, concregation.Name, concregation.Age, concregation.Address, concregation.OrganizerId)
	defer stmt.Close()

	if err != nil {
		return &models.Concregation{}, err
	}

	concregationId, _ := concregationResponse.LastInsertId()
	concregation.ID = concregationId

	return concregation, err
}

func (m *ConcregationRepository) GetByOrganizerId(ctx context.Context, organizerId int64) ([]*models.Concregation, error) {
	query := "Select * From concregation where organizer_id=?"

	return m.fetch(ctx, query, organizerId)
}

func (m *ConcregationRepository) Update(ctx context.Context, p *models.Concregation) (*models.Concregation, error) {
	query := "Update concregation set name=?, age=?, address=? where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.Age,
		p.Address,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *ConcregationRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From concregation Where id=?"

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
