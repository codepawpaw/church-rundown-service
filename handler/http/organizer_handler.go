package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	driver "../../driver"
	models "../../models"
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

func (organizerHandler *OrganizerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"]
	name := r.URL.Query()["name"]

	idParam := ""
	nameParam := ""

	if len(id) > 0 {
		idParam = id[0]
	}

	if len(name) > 0 {
		nameParam = name[0]
	}

	payload, _ := organizerHandler.repository.GetAll(r.Context(), 100, idParam, nameParam)

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a New Organizer
func (organizerHandler *OrganizerHandler) CreateOrganizer(w http.ResponseWriter, r *http.Request) {
	organizer := models.Organizer{}
	json.NewDecoder(r.Body).Decode(&organizer)

	newId, err := organizerHandler.repository.Create(r.Context(), &organizer)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500")
		return
	}

	fmt.Println(newId)

	respondwithJSON(w, http.StatusCreated, "Created")
}

func (organizerHandler *OrganizerHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	organizer := models.Organizer{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&organizer)
	payload, err := organizerHandler.repository.Update(r.Context(), &organizer)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a post details
func (organizerHandler *OrganizerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := organizerHandler.repository.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a post
func (organizerHandler *OrganizerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := organizerHandler.repository.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}
