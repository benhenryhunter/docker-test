package alert

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/dickmanben/rebel-api/helper"
	"github.com/gorilla/mux"
)

// All Send to All Handler
func All(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	alert := Alert{}
	err := helper.Marshaller(r, &alert)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			panic(err)
		}
	}
	err = SendAlertAll(alert)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode("Success"); err != nil {
		panic(err)
	}
}

// GetAll returns all Alert
func GetAll(w http.ResponseWriter, r *http.Request) {

	alerts, err := getAlerts()
	if err != nil {
		if err = json.NewEncoder(w).Encode("Error retreiving alerts: " + err.Error()); err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&alerts); err != nil {
		panic(err)
	}
}

// Get returns singular alert
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	AlertID := bson.ObjectIdHex(vars["AlertID"])

	alert, err := getAlert(AlertID)
	if err != nil {
		if err = json.NewEncoder(w).Encode("Error retreiving alert: " + err.Error()); err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&alert); err != nil {
		panic(err)
	}
}
