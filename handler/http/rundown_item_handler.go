package handler

import (
	"encoding/json"
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

func (rundownItemHandler *RundownItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	rundownItem := models.RundownItem{}
	json.NewDecoder(r.Body).Decode(&rundownItem)

	createdRundownItem, err := rundownItemHandler.repository.Create(r.Context(), &rundownItem)

	rundownResponse, _ := json.Marshal(createdRundownItem)

	response := construct(rundownResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (rundownItemHandler *RundownItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	rundownItem := models.RundownItem{}
	json.NewDecoder(r.Body).Decode(&rundownItem)

	updatedRundownItem, err := rundownItemHandler.repository.Update(r.Context(), &rundownItem)

	rundownResponse, _ := json.Marshal(updatedRundownItem)

	response := construct(rundownResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (rundownItemHandler *RundownItemHandler) GetByRundownId(w http.ResponseWriter, r *http.Request) {
	rundownId, _ := strconv.Atoi(chi.URLParam(r, "rundownId"))
	selectedRundownItem, err := rundownItemHandler.repository.GetByRundownId(r.Context(), int64(rundownId))

	rundownResponse, _ := json.Marshal(selectedRundownItem)

	response := construct(rundownResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (rundownItemHandler *RundownItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := rundownItemHandler.repository.Delete(r.Context(), int64(id))

	rundownResponse, _ := json.Marshal("")

	response := construct(rundownResponse, err)

	respondwithJSON(w, response.Status, response)
}
