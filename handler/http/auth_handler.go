package handler

import (
	"encoding/json"
	"net/http"

	driver "../../driver"
	models "../../models"
	repository "../../repository"
	accountRepository "../../repository/account"
	jwtService "../../service/jwt"
	"github.com/dgrijalva/jwt-go"
)

func InitAuthHandler(db *driver.DB, jwtServiceObj *jwtService.JwtService) *AuthHandler {
	return &AuthHandler{
		accountRepository: accountRepository.InitAccountRepository(db.SQL),
		jwtService:        jwtServiceObj,
	}
}

type AuthHandler struct {
	accountRepository repository.AccountRepository
	jwtService        *jwtService.JwtService
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (authHandler *AuthHandler) Login(response http.ResponseWriter, request *http.Request) {
	account := models.Account{}
	json.NewDecoder(request.Body).Decode(&account)

	payload, err := authHandler.accountRepository.GetByUsernameAndPassword(request.Context(), string(account.Username), string(account.Password))

	if payload == nil {
		respondWithError(response, http.StatusUnauthorized, "Unauthorized")
	}

	if err != nil {
		respondWithError(response, http.StatusUnauthorized, "Unauthorized")
	}

	generatedToken := authHandler.jwtService.Encode(account.Username)

	respondwithJSON(response, http.StatusOK, generatedToken)
}

func (authHandler *AuthHandler) Register(response http.ResponseWriter, request *http.Request) {
	account := models.Account{}
	json.NewDecoder(request.Body).Decode(&account)

	payload, err := authHandler.accountRepository.GetByUsername(request.Context(), string(account.Username))

	if payload != nil {
		respondWithError(response, http.StatusUnavailableForLegalReasons, "Username Already Exists")
	}

	if err != nil {
	}

	authHandler.accountRepository.Create(request.Context(), &account)

	generatedToken := authHandler.jwtService.Encode(account.Username)

	respondwithJSON(response, http.StatusOK, generatedToken)
}
