package user

import (
	"GoForBeginner/internal/mytypes"
	"GoForBeginner/internal/repository"
	"GoForBeginner/internal/utils"
	"github.com/go-chi/chi/v5"
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

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload mytypes.RegisterUserPayload
	err := utils.ParseJson(r, payload)
	if err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, err)
		return
	}

}
