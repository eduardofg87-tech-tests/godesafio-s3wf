package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/user"
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/entity"
	"github.com/GuilhermeCaruso/bellt"
)

func userIndex(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading users"
		var data []*entity.User
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = service.FindAll()
		default:
			data, err = service.Search(name)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func userAdd(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding user"
		var b *entity.User
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		b.ID, err = service.Store(b)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(b); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func userFind(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading user"
		vars := mux.Vars(r)
		id := vars["id"]
		data, err := service.Find(entity.StringToID(id))
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func userDelete(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing user"
		vars := mux.Vars(r)
		id := vars["id"]
		err := service.Delete(entity.StringToID(id))
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeUserHandlers make url handlers
func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.UseCase) {
	r.Handle("/v1/ping", n.With(
		negroni.Wrap(userIndex(service)),
	)).Methods("GET", "OPTIONS").Name("userIndex")

	r.Handle("/v1/user", n.With(
		negroni.Wrap(userAdd(service)),
	)).Methods("POST", "OPTIONS").Name("userAdd")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(userFind(service)),
	)).Methods("GET", "OPTIONS").Name("userFind")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(userDelete(service)),
	)).Methods("DELETE", "OPTIONS").Name("userDelete")
}
