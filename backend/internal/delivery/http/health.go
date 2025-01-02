package http

import (
	"encoding/json"
	"net/http"
)

type health struct {
	Descr string `json:"descr"`
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(health{
		Descr: "Сервер работает корректно!",
	})
}
