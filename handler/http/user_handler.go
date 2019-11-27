package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	driver "../../driver"
	models "../../models"
	repository "../../repository"
	userRepository "../../repository/user"
	"github.com/go-chi/chi"
)

func InitUserHandler(db *driver.DB) *UserHandler {
	return &UserHandler{
		repository: userRepository.InitUserRepository(db.SQL),
	}
}

type UserHandler struct {
	repository repository.UserRepository
}

func (userHandler *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	payload, _ := userHandler.repository.GetAll(r.Context(), 100)

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a New Organizer
func (userHandler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)

	newId, err := userHandler.repository.Create(r.Context(), &user)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500")
		return
	}

	type Data struct {
		InsertedId int64
	}

	data := Data{
		InsertedId: newId,
	}

	respondwithJSON(w, http.StatusCreated, data)
}

func (userHandler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user := models.User{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&user)
	payload, err := userHandler.repository.Update(r.Context(), &user)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a post details
func (userHandler *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := userHandler.repository.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a post
func (userHandler *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := userHandler.repository.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}
