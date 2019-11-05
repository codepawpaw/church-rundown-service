package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	driver "../../driver"
	models "../../models"
	repository "../../repository"
	rundownRepository "../../repository/rundown"
	"github.com/go-chi/chi"
)

func InitRundownHandler(db *driver.DB) *RundownHandler {
	return &RundownHandler{
		repository: rundownRepository.InitRundownRepository(db.SQL),
	}
}

type RundownHandler struct {
	repository repository.RundownRepository
}

func (rundownHandler *RundownHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	payload, _ := rundownHandler.repository.GetAll(r.Context(), 100)

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a New Organizer
func (rundownHandler *RundownHandler) Create(w http.ResponseWriter, r *http.Request) {
	rundown := models.Rundown{}
	json.NewDecoder(r.Body).Decode(&rundown)

	newId, err := rundownHandler.repository.Create(r.Context(), &rundown)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500")
		return
	}

	fmt.Println(newId)

	respondwithJSON(w, http.StatusCreated, "Created")
}

func (rundownHandler *RundownHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	rundown := models.Rundown{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&rundown)
	payload, err := rundownHandler.repository.Update(r.Context(), &rundown)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a post details
func (rundownHandler *RundownHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := rundownHandler.repository.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (rundownHandler *RundownHandler) GetByOrganizerAndId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	organizerId, _ := strconv.Atoi(chi.URLParam(r, "organizerId"))

	payload, err := rundownHandler.repository.GetByOrganizerAndId(r.Context(), int64(id), int64(organizerId))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (rundownHandler *RundownHandler) GetByOrganizerIdAndDate(w http.ResponseWriter, r *http.Request) {
	organizerId, _ := strconv.Atoi(chi.URLParam(r, "organizerId"))
	startDate, _ := strconv.Atoi(chi.URLParam(r, "startDate"))
	endDate, _ := strconv.Atoi(chi.URLParam(r, "endDate"))

	payload, err := rundownHandler.repository.GetByOrganizerIdAndDate(r.Context(), int64(organizerId), string(startDate), string(endDate))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a post details
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

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a post
func (rundownHandler *RundownHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := rundownHandler.repository.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}
