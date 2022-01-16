package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mohamedelkashif/store-location-service/db"
	"github.com/mohamedelkashif/store-location-service/model"
	"io/ioutil"
	"net/http"
)

func GetStores(w http.ResponseWriter, _ *http.Request) {
	stores := db.FindAll()

	bytes, err := json.Marshal(stores)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

func SaveStore(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	store := new(model.Store)
	err = json.Unmarshal(body, store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	db.Save(store.Name, store)
	w.Header().Set("Location", r.URL.Path+"/"+store.StoreId)
	w.WriteHeader(http.StatusCreated)
}

func GetStoresByCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country_code := vars["country"]
	max := vars["max"]

	store := db.FindAllByCountry(country_code, max)
	// if !ok {
	// 	http.NotFound(w, r)
	// 	return
	// }
	bytes, err := json.Marshal(store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, bytes)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
