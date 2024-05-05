package user

import (
	"fmt"
	"net/http"

	"github.com/atiwat-r/golang-backend/service/auth"
	"github.com/atiwat-r/golang-backend/types"
	"github.com/atiwat-r/golang-backend/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}


func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	// Obtain JSON from request, store it in payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// Check if user exist
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user [%s] already exist", payload.Email))
		return
	}

	// Hash the password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// If user doesn't exist, create the user
	err = h.store.CreateUser(types.User{
		Firstname: payload.Firstname,
		Lastname: payload.Lastname,
		Email: payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}