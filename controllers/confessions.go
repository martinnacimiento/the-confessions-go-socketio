package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
)

var confessions []models.Confession

// IndexConfession list all confessions
func IndexConfession(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(confessions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// StoreConfession create a new confession
func StoreConfession(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var confession models.Confession
	if err := decoder.Decode(&confession); err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	confessions = append(confessions, confession)
	fmt.Fprintf(w, "Created!")
}

// UpdateConfession update a confession
func UpdateConfession(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user updated")
}

// DeleteConfession delete a confession
func DeleteConfession(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user deleted")
}

type metaData interface{}
