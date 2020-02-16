package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	driver "../../driver"
	models "../../models"
	repository "../../repository"
	serviceScheduleRepository "../../repository/service_schedule"
	"github.com/go-chi/chi"
)

func InitServiceScheduleHandler(db *driver.DB) *ServiceScheduleHandler {
	return &ServiceScheduleHandler{
		serviceScheduleRepository: serviceScheduleRepository.InitServiceScheduleRepository(db.SQL),
	}
}

type ServiceScheduleHandler struct {
	serviceScheduleRepository repository.ServiceScheduleRepository
}

func (serviceScheduleHandler *ServiceScheduleHandler) Create(w http.ResponseWriter, r *http.Request) {
	serviceSchedule := models.ServiceSchedule{}
	json.NewDecoder(r.Body).Decode(&serviceSchedule)

	serviceScheduleData, err := serviceScheduleHandler.serviceScheduleRepository.Create(r.Context(), &serviceSchedule)

	serviceScheduleResponse, _ := json.Marshal(serviceScheduleData)

	response := construct(serviceScheduleResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (serviceScheduleHandler *ServiceScheduleHandler) Update(w http.ResponseWriter, r *http.Request) {
	serviceSchedule := models.ServiceSchedule{}
	json.NewDecoder(r.Body).Decode(&serviceSchedule)
	payload, err := serviceScheduleHandler.serviceScheduleRepository.Update(r.Context(), &serviceSchedule)

	serviceScheduleResponse, _ := json.Marshal(payload)

	response := construct(serviceScheduleResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (serviceScheduleHandler *ServiceScheduleHandler) GetByOrganizerId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "organizerId"))

	payload, err := serviceScheduleHandler.serviceScheduleRepository.GetByOrganizerId(r.Context(), int64(id))

	serviceScheduleResponse, _ := json.Marshal(payload)

	response := construct(serviceScheduleResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (serviceScheduleHandler *ServiceScheduleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	payload, err := serviceScheduleHandler.serviceScheduleRepository.Delete(r.Context(), int64(id))

	serviceScheduleResponse, _ := json.Marshal(payload)

	response := construct(serviceScheduleResponse, err)

	respondwithJSON(w, response.Status, response)

}
