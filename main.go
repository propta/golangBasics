package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	ZipCode string `json:"zip_code"`
	City    string `json:"city"`
}

func main() {
	http.HandleFunc("/users", getAllUsers)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		return
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{"Dennis", "20001", "Nairobi"},
		{"Kiprop", "390210", "Kericho"},
		{"Silot", "3271", "Mombasa"},
		{"Reynold", "2021001", "Kisumu"},
	}
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}
