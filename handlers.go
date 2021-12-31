package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the City Database!")

}

func CityList(w http.ResponseWriter, r *http.Request) {
	jsonCities := dbCityList()

	w.Header().Set("Content-Type", "application/jason")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCities)
}

func CityDisplay(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityId, _ := strconv.Atoi(vars["id"])

	jsonCity := dbCityDisplay(cityId)

	w.Header().Set("Content-Type", "application/jason")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCity)

}

func CityAdd(w http.ResponseWriter, r *http.Request) {
	var city City

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &city); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422) // can't process!
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	addResult := dbCityAdd(city)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(addResult)
}

func CityDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityId, _ := strconv.ParseInt(vars["id"], 10, 64)

	deleteResult := dbCityDelete(cityId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(deleteResult)
}
