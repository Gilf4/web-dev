package user

import (
	"GoForBeginner/internal/db/models"
	"GoForBeginner/internal/mytypes"
	"GoForBeginner/internal/repository"
	"GoForBeginner/internal/utils"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Handler struct {
	userRepo repository.UserStoreRepository
}

func NewHandler(userRepo repository.UserStoreRepository) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}

func (h *Handler) RegisterRoutes(router chi.Router) {
	router.Post("/login", h.handleLogin)
	router.Post("/register", h.handleRegister)

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload mytypes.LoginUserPayload
	err := utils.ParseJson(r, &payload)
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, err)
		return
	}

	user, err := h.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, err)
		return
	}
	if user == nil {
		utils.WriteJsonError(w, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	if !utils.VerifyPassword(payload.Password, user.Password) {
		utils.WriteJsonError(w, http.StatusUnauthorized, fmt.Errorf("incorrect password"))
		return
	}

	utils.WriteJson(w, http.StatusOK, payload)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload mytypes.RegisterUserPayload
	err := utils.ParseJson(r, &payload)
	if err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, err)
		return
	}

	// validate
	err = utils.Validate.Struct(payload)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJsonError(w, http.StatusBadRequest, fmt.Errorf("validation error: %v", errors))
		return
	}

	// check if the user exists
	existingUser, err := h.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError,
			fmt.Errorf("failed to check user existence: %w", err))
		return
	}
	if existingUser != nil {
		utils.WriteJsonError(w, http.StatusBadRequest,
			fmt.Errorf("user with email: %v already exists", payload.Email))
		return
	}

	// hashing password
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, err)
		return
	}

	// create user
	err = h.userRepo.CreateUser(models.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
		Nickname:  payload.Nickname,
	})
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, nil)
}
