package mongo

import (
	"encoding/json"
	"net/http"
)

func GetAllScoresHandler(w http.ResponseWriter, r *http.Request) {
	res, err := GetAllScores()
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(res)
}

func GetOneByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := GetOneById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(res)

}
