package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	driver "../../driver"
	models "../../models"
	repository "../../repository"
	accountRepository "../../repository/account"
	"github.com/go-chi/chi"
)

func InitAccountHandler(db *driver.DB) *AccountHandler {
	return &AccountHandler{
		repository: accountRepository.InitAccountRepository(db.SQL),
	}
}

type AccountHandler struct {
	repository repository.AccountRepository
}

func (accountHandler *AccountHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	payload, _ := accountHandler.repository.GetAll(r.Context(), 100)

	respondwithJSON(w, http.StatusOK, payload)
}

func (accountHandler *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	account := models.Account{}
	json.NewDecoder(r.Body).Decode(&account)

	newId, err := accountHandler.repository.Create(r.Context(), &account)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500")
		return
	}

	fmt.Println(newId)

	respondwithJSON(w, http.StatusCreated, "Created")
}

func (accountHandler *AccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	account := models.Account{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&account)
	payload, err := accountHandler.repository.Update(r.Context(), &account)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (accountHandler *AccountHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := accountHandler.repository.GetByID(r.Context(), int64(id))

	accountResponse, _ := json.Marshal(payload)

	response := construct(accountResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (accountHandler *AccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := accountHandler.repository.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
