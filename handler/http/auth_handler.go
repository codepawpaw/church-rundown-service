package handler

import (
	"encoding/json"
	"net/http"

	driver "../../driver"
	dto "../../dto"
	models "../../models"
	repository "../../repository"
	accountRepository "../../repository/account"
	authRepository "../../repository/auth"
	organizerRepository "../../repository/organizer"
	userRepository "../../repository/user"
	jwtService "../../service/jwt"
	"github.com/dgrijalva/jwt-go"
)

func InitAuthHandler(db *driver.DB, jwtServiceObj *jwtService.JwtService) *AuthHandler {
	return &AuthHandler{
		accountRepository:   accountRepository.InitAccountRepository(db.SQL),
		authRepository:      authRepository.InitAuthRepository(db.SQL),
		userRepository:      userRepository.InitUserRepository(db.SQL),
		organizerRepository: organizerRepository.InitOrganizerRepository(db.SQL),
		jwtService:          jwtServiceObj,
	}
}

type AuthHandler struct {
	accountRepository   repository.AccountRepository
	authRepository      repository.AuthRepository
	userRepository      repository.UserRepository
	organizerRepository repository.OrganizerRepository
	jwtService          *jwtService.JwtService
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (authHandler *AuthHandler) Login(response http.ResponseWriter, request *http.Request) {
	account := models.Account{}
	json.NewDecoder(request.Body).Decode(&account)

	accountResponse, err := authHandler.accountRepository.GetByUsernameAndPassword(request.Context(), string(account.Username), string(account.Password))

	if accountResponse == nil {
		respondWithError(response, http.StatusUnauthorized, "Unauthorized")
	}

	if err != nil {
		respondWithError(response, http.StatusUnauthorized, "Unauthorized")
	}

	userReponse, err := authHandler.userRepository.GetByID(request.Context(), accountResponse.UserId)
	organizerResponse, _ := authHandler.organizerRepository.GetByID(request.Context(), userReponse.OrganizerId)

	generatedToken := authHandler.jwtService.Encode(accountResponse.Username)

	authResponse := dto.Auth{
		Account:   accountResponse,
		User:      userReponse,
		Organizer: organizerResponse,
		Token:     generatedToken,
	}

	respondwithJSON(response, http.StatusOK, authResponse)
}

func (authHandler *AuthHandler) Register(response http.ResponseWriter, request *http.Request) {
	authModel := models.Auth{}
	json.NewDecoder(request.Body).Decode(&authModel)

	organizer := authModel.Organizer
	user := authModel.User
	account := authModel.Account

	authResponse := authHandler.authRepository.Create(request.Context(), &organizer, &user, &account)

	generatedToken := authHandler.jwtService.Encode(account.Username)

	authResponse.Token = generatedToken

	respondwithJSON(response, http.StatusOK, authResponse)
}
