package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	ZipCode string `json:"zip_code"`
	City    string `json:"city"`
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{"Dennis", "20001", "Nairobi"},
		{"Kiprop", "390210", "Kericho"},
		{"Silot", "3271", "Mombasa"},
		{"Reynold", "2021001", "Kisumu"},
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			return
		}
	} else {
		w.Header().Add("Content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(users)
		if err != nil {
			return
		}
	}

}
