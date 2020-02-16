package service_schedule

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitServiceScheduleRepository(Connection *sql.DB) pRepo.ServiceScheduleRepository {
	return &ServiceScheduleRepository{
		Connection: Connection,
	}
}

type ServiceScheduleRepository struct {
	Connection *sql.DB
}

func (o *ServiceScheduleRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.ServiceSchedule, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.ServiceSchedule, 0)
	for rows.Next() {
		data := new(models.ServiceSchedule)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Text,
			&data.Date,
			&data.OrganizerId,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (o *ServiceScheduleRepository) Create(ctx context.Context, serviceSchedule *models.ServiceSchedule) (*models.ServiceSchedule, error) {
	query := "Insert service_schedule SET name=?, text=?, date=?, organizer_id=?"

	stmt, err := o.Connection.PrepareContext(ctx, query)
	if err != nil {
		return &models.ServiceSchedule{}, err
	}

	serviceScheduleResponse, err := stmt.ExecContext(ctx, serviceSchedule.Name, serviceSchedule.Text, serviceSchedule.Date, serviceSchedule.OrganizerId)
	defer stmt.Close()

	if err != nil {
		return &models.ServiceSchedule{}, err
	}

	serviceScheduleId, _ := serviceScheduleResponse.LastInsertId()
	serviceSchedule.ID = serviceScheduleId

	return serviceSchedule, err
}

func (m *ServiceScheduleRepository) GetByOrganizerId(ctx context.Context, serviceScheduleId int64) ([]*models.ServiceSchedule, error) {
	query := "Select * From service_schedule where organizer_id=?"

	return m.fetch(ctx, query, serviceScheduleId)
}

func (m *ServiceScheduleRepository) Update(ctx context.Context, p *models.ServiceSchedule) (*models.ServiceSchedule, error) {
	query := "Update service_schedule set name=?, text=?, date=? where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.Text,
		p.Date,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *ServiceScheduleRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From service_schedule Where id=?"

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
