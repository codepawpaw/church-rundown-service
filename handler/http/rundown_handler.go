package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	driver "../../driver"
	dto "../../dto"
	models "../../models"
	repository "../../repository"
	rundownRepository "../../repository/rundown"
	rundownItemRepository "../../repository/rundown_item"
	"github.com/go-chi/chi"
)

func InitRundownHandler(db *driver.DB) *RundownHandler {
	return &RundownHandler{
		repository:            rundownRepository.InitRundownRepository(db.SQL),
		rundownItemRepository: rundownItemRepository.InitRundownItemRepository(db.SQL),
	}
}

type RundownHandler struct {
	repository            repository.RundownRepository
	rundownItemRepository repository.RundownItemRepository
}

type RundownHttpResponse struct {
	Data     string            `json:"data"`
	Response *dto.HttpResponse `json:"response"`
}

func (rundownHandler *RundownHandler) Create(w http.ResponseWriter, r *http.Request) {
	rundown := models.Rundown{}
	json.NewDecoder(r.Body).Decode(&rundown)

	rundownData, err := rundownHandler.repository.Create(r.Context(), &rundown)

	rundownResponse, _ := json.Marshal(rundownData)

	response := construct(rundownResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (rundownHandler *RundownHandler) Update(w http.ResponseWriter, r *http.Request) {
	rundown := models.Rundown{}
	json.NewDecoder(r.Body).Decode(&rundown)
	payload, err := rundownHandler.repository.Update(r.Context(), &rundown)

	rundownResponse, _ := json.Marshal(payload)

	response := construct(rundownResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (rundownHandler *RundownHandler) GetByOrganizerIdAndDate(w http.ResponseWriter, r *http.Request) {
	organizerId, _ := strconv.Atoi(chi.URLParam(r, "organizerId"))
	startDate, _ := strconv.Atoi(chi.URLParam(r, "startDate"))
	endDate, _ := strconv.Atoi(chi.URLParam(r, "endDate"))

	payload, err := rundownHandler.repository.GetByOrganizerIdAndDate(r.Context(), int64(organizerId), string(startDate), string(endDate))

	rundownResponse, _ := json.Marshal(payload)

	response := construct(rundownResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (rundownHandler *RundownHandler) GetByOrganizerId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "organizerId"))

	startDate := r.URL.Query()["startDate"]
	endDate := r.URL.Query()["endDate"]

	startDateParam := ""
	endDateParam := ""

	if len(startDate) > 0 {
		startDateParam = startDate[0]
	}

	if len(endDate) > 0 {
		endDateParam = endDate[0]
	}

	payload, err := rundownHandler.repository.GetByOrganizerId(r.Context(), int64(id), startDateParam, endDateParam)

	rundownResponse, _ := json.Marshal(payload)

	response := construct(rundownResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (rundownHandler *RundownHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	isDeleteSuccess, _ := rundownHandler.rundownItemRepository.DeleteByRundownId(r.Context(), int64(id))

	if isDeleteSuccess == false {
		respondwithJSON(w, http.StatusInternalServerError, "Failed to delete rundown item")
	}

	payload, err := rundownHandler.repository.Delete(r.Context(), int64(id))

	rundownResponse, _ := json.Marshal(payload)

	response := construct(rundownResponse, err)

	respondwithJSON(w, response.Status, response)

}
