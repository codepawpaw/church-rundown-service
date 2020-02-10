package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	driver "../../driver"
	repository "../../repository"
	organizerRepository "../../repository/organizer"
	"github.com/go-chi/chi"
)

func InitOrganizerHandler(db *driver.DB) *OrganizerHandler {
	return &OrganizerHandler{
		repository: organizerRepository.InitOrganizerRepository(db.SQL),
	}
}

type OrganizerHandler struct {
	repository repository.OrganizerRepository
}

func (organizerHandler *OrganizerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := organizerHandler.repository.GetByID(r.Context(), int64(id))

	organizerResponse, _ := json.Marshal(payload)

	response := construct(organizerResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (organizerHandler *OrganizerHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	payload, err := organizerHandler.repository.GetByName(r.Context(), string(name))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (organizerHandler *OrganizerHandler) GetByCity(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")
	payload, err := organizerHandler.repository.GetByCity(r.Context(), string(city))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (organizerHandler *OrganizerHandler) GetByProvince(w http.ResponseWriter, r *http.Request) {
	province := chi.URLParam(r, "province")
	payload, err := organizerHandler.repository.GetByProvince(r.Context(), string(province))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}
