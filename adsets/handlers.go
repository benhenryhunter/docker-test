package adsets

import (
	"encoding/json"
	"net/http"

	"github.com/dickmanben/rebel-api/helper"

	"github.com/dickmanben/rebel-api/alert"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func PassThru(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	AdSetID := bson.ObjectIdHex(vars["AdSetID"])

	adset, err := getPassThru(AdSetID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode("Error getting pass thru: " + err.Error()); err != nil {
			alert.SendAlertAll(alert.Alert{})
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write(adset)
}

func CreateAdset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adset := AdSet{}
	err := helper.Marshaller(r, &adset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode("Error getting pass thru: " + err.Error()); err != nil {
			alert.SendAlertAll(alert.Alert{})
		}
	}

	adset, err = createAdset(adset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode("Error getting pass thru: " + err.Error()); err != nil {
			alert.SendAlertAll(alert.Alert{})
		}
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&adset); err != nil {
		alert.SendAlertAll(alert.Alert{})
	}
}

func UpdateAdset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adset := AdSet{}
	err := helper.Marshaller(r, &adset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode("Error getting pass thru: " + err.Error()); err != nil {
			alert.SendAlertAll(alert.Alert{})
		}
	}

	vars := mux.Vars(r)
	AdSetID := bson.ObjectIdHex(vars["AdSetID"])

	adset, err = updateAdset(adset, AdSetID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode("Error getting pass thru: " + err.Error()); err != nil {
			alert.SendAlertAll(alert.Alert{})
		}
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&adset); err != nil {
		alert.SendAlertAll(alert.Alert{})
	}
}

func GetAdset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	AdSetID := bson.ObjectIdHex(vars["AdSetID"])

	adset, err := getAdset(AdSetID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode("Error getting pass thru: " + err.Error()); err != nil {
			alert.SendAlertAll(alert.Alert{})
		}
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&adset); err != nil {
		alert.SendAlertAll(alert.Alert{})
	}
}

func GetAdsets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adsets, err := getAdsets()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err = json.NewEncoder(w).Encode("Error getting pass thru: " + err.Error()); err != nil {
			alert.SendAlertAll(alert.Alert{})
		}
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(&adsets); err != nil {
		alert.SendAlertAll(alert.Alert{})
	}
}
