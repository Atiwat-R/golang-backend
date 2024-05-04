package user

import (
	"net/http"

	"github.com/atiwat-r/golang-backend/types"
	"github.com/atiwat-r/golang-backend/utils"
	"github.com/gorilla/mux"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
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

}