package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/verzano/go-home/pkg/gome/model"
	"log"
	"net/http"
)

var devicesMap = map[string]model.Device{
	"dev1": {
		Id:     "dev1",
		Status: "healthy",
	},
	"dev2": {
		Id:     "dev2",
		Status: "unhealthy",
	},
}

func getDevices(w http.ResponseWriter, r *http.Request) {
	devices := make([]model.Device, len(devicesMap))

	for _, d := range devicesMap {
		devices = append(devices, d)
	}

	response, _ := json.Marshal(devices)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func getDeviceById(w http.ResponseWriter, r *http.Request) {
	device := devicesMap[mux.Vars(r)["id"]]

	response, _ := json.Marshal(device)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Init() {
	router := mux.NewRouter()
	router.HandleFunc("/devices", getDevices).Methods("GET")
	router.HandleFunc("/devices/{id}", getDeviceById).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}
