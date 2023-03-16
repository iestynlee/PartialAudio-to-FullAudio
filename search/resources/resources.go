package resources

import (
	"encoding/json"
	"net/http"
	"search/service"

	"github.com/gorilla/mux"
)

func searchAudio(w http.ResponseWriter, r *http.Request) {
	a := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&a); err == nil {
		/* Gets audio from the POST request */
		if audio, ok := a["Audio"].(string); ok {
			if title, err := service.Service(audio); err == nil {
				/* Writing the audio id ouput and encoding to the response */
				u := map[string]interface{}{"id": title}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(u)
				return
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
	/* controller */
	r.HandleFunc("/search", searchAudio).Methods("POST")
	return r
}
