package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"backend/internal/domain"
	"backend/internal/usecase"

	"github.com/go-chi/chi"
)

type AdminHandler struct {
	adminUsecase usecase.AdminUsecase
}

func NewAdminHandler(adminUsecase usecase.AdminUsecase) *AdminHandler {
	return &AdminHandler{adminUsecase}
}

func (h *AdminHandler) GetAllAdmins(w http.ResponseWriter, r *http.Request) {
	admins, err := h.adminUsecase.GetAllAdmins()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(admins)
}

func (h *AdminHandler) GetAdminByName(w http.ResponseWriter, r *http.Request) {
	var admin domain.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		fmt.Println("body:", r.Body)
		fmt.Println("admin: ", admin, ", err:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("admin-requests: name -", admin.Username, "password -", admin.Password)

	admin, err := h.adminUsecase.GetAdminByName(admin.Username, admin.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(admin)
}

func (h *AdminHandler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var admin domain.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.adminUsecase.CreateAdmin(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(admin)
}

func (h *AdminHandler) UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var admin domain.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	admin.ID = uint(id)
	if err := h.adminUsecase.UpdateAdmin(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(admin)
}

func (h *AdminHandler) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.adminUsecase.DeleteAdmin(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
