package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	driver "../../driver"
	models "../../models"
	repository "../../repository"
	deviceInventoryRepository "../../repository/device_inventory"
	"github.com/go-chi/chi"
)

func InitDeviceInventoryHandler(db *driver.DB) *DeviceInventoryHandler {
	return &DeviceInventoryHandler{
		deviceInventoryRepository: deviceInventoryRepository.InitDeviceInventoryRepository(db.SQL),
	}
}

type DeviceInventoryHandler struct {
	deviceInventoryRepository repository.DeviceInventoryRepository
}

func (deviceInventoryHandler *DeviceInventoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	deviceInventory := models.DeviceInventory{}
	json.NewDecoder(r.Body).Decode(&deviceInventory)

	deviceInventoryData, err := deviceInventoryHandler.deviceInventoryRepository.Create(r.Context(), &deviceInventory)

	deviceInventoryResponse, _ := json.Marshal(deviceInventoryData)

	response := construct(deviceInventoryResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (deviceInventoryHandler *DeviceInventoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	deviceInventory := models.DeviceInventory{}
	json.NewDecoder(r.Body).Decode(&deviceInventory)
	payload, err := deviceInventoryHandler.deviceInventoryRepository.Update(r.Context(), &deviceInventory)

	deviceInventoryResponse, _ := json.Marshal(payload)

	response := construct(deviceInventoryResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (deviceInventoryHandler *DeviceInventoryHandler) GetByOrganizerId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "organizerId"))

	payload, err := deviceInventoryHandler.deviceInventoryRepository.GetByOrganizerId(r.Context(), int64(id))

	deviceInventoryResponse, _ := json.Marshal(payload)

	response := construct(deviceInventoryResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (deviceInventoryHandler *DeviceInventoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	payload, err := deviceInventoryHandler.deviceInventoryRepository.Delete(r.Context(), int64(id))

	deviceInventoryResponse, _ := json.Marshal(payload)

	response := construct(deviceInventoryResponse, err)

	respondwithJSON(w, response.Status, response)

}
