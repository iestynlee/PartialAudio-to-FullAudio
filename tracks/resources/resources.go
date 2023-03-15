package resources

import (
	"encoding/json"
	"net/http"
	"tracks/repository"

	"github.com/gorilla/mux"
)

func updateCell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var c repository.Cell
	if err := json.NewDecoder(r.Body).Decode(&c); err == nil {
		if id == c.Id {
			if n := repository.Update(c); n > 0 {
				w.WriteHeader(204) /* No Content */
			} else if n := repository.Insert(c); n > 0 {
				w.WriteHeader(201) /* Created */
			} else {
				w.WriteHeader(500) /* Internal Server Error */
			}
		} else {
			w.WriteHeader(400) /* Bad Request */
		}
	} else {
		w.WriteHeader(400) /* Bad Request */
	}
}

func readCell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if c, n := repository.Read(id); n > 0 {
		d := repository.Cell{Id: c.Id, Audio: c.Audio}
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(d)
	} else if n == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func deleteCell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete the cell with the specified ID
	if n := repository.Delete(id); n > 0 {
		w.WriteHeader(204) /* No Content */
	} else if n == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func allCell(w http.ResponseWriter, r *http.Request) {
	if c, n := repository.All(); n > 0 {
		output := make([]string, len(c))
		for i, cell := range c {
			output[len(c)-1-i] = cell.Id
		}
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(output)
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* Store */
	r.HandleFunc("/tracks/{id}", updateCell).Methods("PUT")
	/* Delete */
	r.HandleFunc("/tracks/{id}", deleteCell).Methods("DELETE")
	/* Document */
	r.HandleFunc("/tracks/{id}", readCell).Methods("GET")
	/* Read all */
	r.HandleFunc("/tracks", allCell).Methods("GET")
	return r
}
