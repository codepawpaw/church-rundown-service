package sector_coordinator

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitSectorCoordinatorRepository(Connection *sql.DB) pRepo.SectorCoordinatorRepository {
	return &SectorCoordinatorRepository{
		Connection: Connection,
	}
}

type SectorCoordinatorRepository struct {
	Connection *sql.DB
}

func (o *SectorCoordinatorRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.SectorCoordinator, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.SectorCoordinator, 0)
	for rows.Next() {
		data := new(models.SectorCoordinator)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.ConcregationId,
			&data.OrganizerId,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (o *SectorCoordinatorRepository) Create(ctx context.Context, sectorCoordinator *models.SectorCoordinator) (*models.SectorCoordinator, error) {
	query := "Insert sector_coordinator SET name=?, concregation_id=?, organizer_id=?"

	stmt, err := o.Connection.PrepareContext(ctx, query)
	if err != nil {
		return &models.SectorCoordinator{}, err
	}

	sectorCoordinatorResponse, err := stmt.ExecContext(ctx, sectorCoordinator.Name, sectorCoordinator.ConcregationId, sectorCoordinator.OrganizerId)
	defer stmt.Close()

	if err != nil {
		return &models.SectorCoordinator{}, err
	}

	sectorCoordinatorId, _ := sectorCoordinatorResponse.LastInsertId()
	sectorCoordinator.ID = sectorCoordinatorId

	return sectorCoordinator, err
}

func (m *SectorCoordinatorRepository) GetByOrganizerId(ctx context.Context, sectorCoordinatorId int64) ([]*models.SectorCoordinator, error) {
	query := "Select * From sector_coordinator where organizer_id=?"

	return m.fetch(ctx, query, sectorCoordinatorId)
}

func (m *SectorCoordinatorRepository) Update(ctx context.Context, p *models.SectorCoordinator) (*models.SectorCoordinator, error) {
	query := "Update sector_coordinator set name=?, concregation_id=? where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.ConcregationId,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *SectorCoordinatorRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From sector_coordinator Where id=?"

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
