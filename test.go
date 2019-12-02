package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//Account is used to hold the data
type Account struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber int    `json:"phonenumber"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
	Status      string `json:"status"`
}

var acc = make(map[string]Account)

func main() {
	initHandler()

}
func initHandler() {

	r := mux.NewRouter()
	r.HandleFunc("/account", createAccount).Methods(http.MethodPost)
	r.HandleFunc("/account", getAccount).Methods(http.MethodGet)
	r.HandleFunc("/account", updateAccount).Methods(http.MethodPut)
	r.HandleFunc("/account", deleteAccount).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", r)

}

func createAccount(w http.ResponseWriter, r *http.Request) {
	var a Account

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return

	}

	if a.ID == "" {

		a.ID = uuid.New().String()
	}
	a.Status = "ACTIVE"
	acc[a.ID] = a
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Account created Successfully")

}

func getAccount(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.URL.Query()["id"]; ok {
		if err := json.NewEncoder(w).Encode(acc[id[0]]); err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	} else {

		if err := json.NewEncoder(w).Encode(acc); err != nil {
			w.Write([]byte(err.Error()))
			return

		}

	}
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.URL.Query()["id"]; ok {

		var a Account

		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		a.ID = id[0]
		acc[id[0]] = a
		a.Status = "ACTIVE"
		w.WriteHeader(http.StatusNoContent)

		return
	}

	w.WriteHeader(http.StatusBadRequest)
	return

}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.URL.Query()["id"]; ok {

		a := acc[id[0]]
		a.Status = "DELETED"

		acc[id[0]] = a
		w.WriteHeader(http.StatusNoContent)

		return
	}

	w.WriteHeader(http.StatusBadRequest)
	return

}
