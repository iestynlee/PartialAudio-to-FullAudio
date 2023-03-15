package resources

import (
	"cooltown/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func collectFullAudio(w http.ResponseWriter, r *http.Request) {
	a := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&a); err == nil {
		if audioPartial, ok := a["Audio"].(string); ok {
			if title, err := service.ServiceAudio(audioPartial); err == nil {
				if audioFull, err := service.ServiceTracks(title); err == nil {
					u := map[string]interface{}{"Audio": audioFull}
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(u)
					return
				} else {
					w.WriteHeader(http.StatusNotFound)
					return
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* Cool Town link */
	r.HandleFunc("/cooltown", collectFullAudio).Methods("POST")
	return r
}
