package device_inventory

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitDeviceInventoryRepository(Connection *sql.DB) pRepo.DeviceInventoryRepository {
	return &DeviceInventoryRepository{
		Connection: Connection,
	}
}

type DeviceInventoryRepository struct {
	Connection *sql.DB
}

func (o *DeviceInventoryRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.DeviceInventory, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.DeviceInventory, 0)
	for rows.Next() {
		data := new(models.DeviceInventory)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.PurchaseDate,
			&data.Total,
			&data.OrganizerId,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (o *DeviceInventoryRepository) Create(ctx context.Context, deviceInventory *models.DeviceInventory) (*models.DeviceInventory, error) {
	query := "Insert device_inventory SET name=?, purchase_date=?, total=?, organizer_id=?"

	stmt, err := o.Connection.PrepareContext(ctx, query)
	if err != nil {
		return &models.DeviceInventory{}, err
	}

	deviceInventoryResponse, err := stmt.ExecContext(ctx, deviceInventory.Name, deviceInventory.PurchaseDate, deviceInventory.Total, deviceInventory.OrganizerId)
	defer stmt.Close()

	if err != nil {
		return &models.DeviceInventory{}, err
	}

	deviceInventoryId, _ := deviceInventoryResponse.LastInsertId()
	deviceInventory.ID = deviceInventoryId

	return deviceInventory, err
}

func (m *DeviceInventoryRepository) GetByOrganizerId(ctx context.Context, deviceInventoryId int64) ([]*models.DeviceInventory, error) {
	query := "Select * From device_inventory where organizer_id=?"

	return m.fetch(ctx, query, deviceInventoryId)
}

func (m *DeviceInventoryRepository) Update(ctx context.Context, p *models.DeviceInventory) (*models.DeviceInventory, error) {
	query := "Update device_inventory set name=?, purchase_date=?, total=? where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.PurchaseDate,
		p.Total,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *DeviceInventoryRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From device_inventory Where id=?"

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
