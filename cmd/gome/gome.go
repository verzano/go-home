package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type device struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

var devicesMap = make(map[string]device)

func getAllDevices(w http.ResponseWriter, r *http.Request) {
	devices := make([]device, len(devicesMap))

	for _, device := range devicesMap {
		devices = append(devices, device)
	}

	response, _ := json.Marshal(devices)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func getDeviceById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	response, _ := json.Marshal(devicesMap[id])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {
	devicesMap["dev1"] = device{
		Id:     "dev1",
		Status: "healthy",
	}
	devicesMap["dev2"] = device{
		Id:     "dev2",
		Status: "unhealthy",
	}

	router := mux.NewRouter()
	router.HandleFunc("/devices", getAllDevices).Methods("GET")
	router.HandleFunc("/devices/{id}", getDeviceById).Methods("GET")

	if err := http.ListenAndServe(":8890", router); err != nil {
		log.Fatal(err)
	}
}
