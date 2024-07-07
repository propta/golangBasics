package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		return
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["customer_id"])
}
