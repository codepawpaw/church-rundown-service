package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	driver "../../driver"
	models "../../models"
	repository "../../repository"
	concregationRepository "../../repository/concregation"
	"github.com/go-chi/chi"
)

func InitConcregationHandler(db *driver.DB) *ConcregationHandler {
	return &ConcregationHandler{
		concregationRepository: concregationRepository.InitConcregationRepository(db.SQL),
	}
}

type ConcregationHandler struct {
	concregationRepository repository.ConcregationRepository
}

func (concregationHandler *ConcregationHandler) Create(w http.ResponseWriter, r *http.Request) {
	concregation := models.Concregation{}
	json.NewDecoder(r.Body).Decode(&concregation)

	concregationData, err := concregationHandler.concregationRepository.Create(r.Context(), &concregation)

	concregationResponse, _ := json.Marshal(concregationData)

	response := construct(concregationResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (concregationHandler *ConcregationHandler) Update(w http.ResponseWriter, r *http.Request) {
	concregation := models.Concregation{}
	json.NewDecoder(r.Body).Decode(&concregation)
	payload, err := concregationHandler.concregationRepository.Update(r.Context(), &concregation)

	concregationResponse, _ := json.Marshal(payload)

	response := construct(concregationResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (concregationHandler *ConcregationHandler) GetByOrganizerId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "organizerId"))

	payload, err := concregationHandler.concregationRepository.GetByOrganizerId(r.Context(), int64(id))

	concregationResponse, _ := json.Marshal(payload)

	response := construct(concregationResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (concregationHandler *ConcregationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	payload, err := concregationHandler.concregationRepository.Delete(r.Context(), int64(id))

	concregationResponse, _ := json.Marshal(payload)

	response := construct(concregationResponse, err)

	respondwithJSON(w, response.Status, response)

}
