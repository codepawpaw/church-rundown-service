package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	driver "../../driver"
	models "../../models"
	repository "../../repository"
	rundownItemRepository "../../repository/rundown_item"
	"github.com/go-chi/chi"
)

func InitRundownItemHandler(db *driver.DB) *RundownItemHandler {
	return &RundownItemHandler{
		repository: rundownItemRepository.InitRundownItemRepository(db.SQL),
	}
}

type RundownItemHandler struct {
	repository repository.RundownItemRepository
}

func (rundownItemHandler *RundownItemHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	payload, _ := rundownItemHandler.repository.GetAll(r.Context(), 100)

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a New Organizer
func (rundownItemHandler *RundownItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	rundownItem := models.RundownItem{}
	json.NewDecoder(r.Body).Decode(&rundownItem)

	newId, err := rundownItemHandler.repository.Create(r.Context(), &rundownItem)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500")
		return
	}

	fmt.Println(newId)

	respondwithJSON(w, http.StatusCreated, "Created")
}

func (rundownItemHandler *RundownItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	rundownItem := models.RundownItem{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&rundownItem)
	payload, err := rundownItemHandler.repository.Update(r.Context(), &rundownItem)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a post details
func (rundownItemHandler *RundownItemHandler) GetByRundownId(w http.ResponseWriter, r *http.Request) {
	rundownId, _ := strconv.Atoi(chi.URLParam(r, "rundownId"))
	payload, err := rundownItemHandler.repository.GetByRundownId(r.Context(), int64(rundownId))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a post
func (rundownItemHandler *RundownItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := rundownItemHandler.repository.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}
