package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/internal/domain"
	"backend/internal/usecase"

	"github.com/go-chi/chi"
)

type TireHandler struct {
	tireUsecase usecase.TireUsecase
}

func NewTireHandler(tireUsecase usecase.TireUsecase) *TireHandler {
	return &TireHandler{tireUsecase}
}

func (h *TireHandler) GetAllTires(w http.ResponseWriter, r *http.Request) {
	tires, err := h.tireUsecase.GetAllTires()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tires)
}

func (h *TireHandler) GetTireByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tire, err := h.tireUsecase.GetTireByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(tire)
}

func (h *TireHandler) CreateTire(w http.ResponseWriter, r *http.Request) {
	var tire domain.Tire
	if err := json.NewDecoder(r.Body).Decode(&tire); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.tireUsecase.CreateTire(&tire); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tire)
}

func (h *TireHandler) UpdateTire(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var tire domain.Tire
	if err := json.NewDecoder(r.Body).Decode(&tire); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tire.ID = uint(id)
	if err := h.tireUsecase.UpdateTire(&tire); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tire)
}

func (h *TireHandler) DeleteTire(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.tireUsecase.DeleteTire(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
