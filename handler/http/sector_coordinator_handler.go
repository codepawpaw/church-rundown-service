package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	driver "../../driver"
	models "../../models"
	repository "../../repository"
	sectorCoordinatorRepository "../../repository/sector_coordinator"
	"github.com/go-chi/chi"
)

func InitSectorCoordinatorHandler(db *driver.DB) *SectorCoordinatorHandler {
	return &SectorCoordinatorHandler{
		sectorCoordinatorRepository: sectorCoordinatorRepository.InitSectorCoordinatorRepository(db.SQL),
	}
}

type SectorCoordinatorHandler struct {
	sectorCoordinatorRepository repository.SectorCoordinatorRepository
}

func (sectorCoordinatorHandler *SectorCoordinatorHandler) Create(w http.ResponseWriter, r *http.Request) {
	sectorCoordinator := models.SectorCoordinator{}
	json.NewDecoder(r.Body).Decode(&sectorCoordinator)

	sectorCoordinatorData, err := sectorCoordinatorHandler.sectorCoordinatorRepository.Create(r.Context(), &sectorCoordinator)

	sectorCoordinatorResponse, _ := json.Marshal(sectorCoordinatorData)

	response := construct(sectorCoordinatorResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (sectorCoordinatorHandler *SectorCoordinatorHandler) Update(w http.ResponseWriter, r *http.Request) {
	sectorCoordinator := models.SectorCoordinator{}
	json.NewDecoder(r.Body).Decode(&sectorCoordinator)
	payload, err := sectorCoordinatorHandler.sectorCoordinatorRepository.Update(r.Context(), &sectorCoordinator)

	sectorCoordinatorResponse, _ := json.Marshal(payload)

	response := construct(sectorCoordinatorResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (sectorCoordinatorHandler *SectorCoordinatorHandler) GetByOrganizerId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "organizerId"))

	payload, err := sectorCoordinatorHandler.sectorCoordinatorRepository.GetByOrganizerId(r.Context(), int64(id))

	sectorCoordinatorResponse, _ := json.Marshal(payload)

	response := construct(sectorCoordinatorResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (sectorCoordinatorHandler *SectorCoordinatorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	payload, err := sectorCoordinatorHandler.sectorCoordinatorRepository.Delete(r.Context(), int64(id))

	sectorCoordinatorResponse, _ := json.Marshal(payload)

	response := construct(sectorCoordinatorResponse, err)

	respondwithJSON(w, response.Status, response)

}
