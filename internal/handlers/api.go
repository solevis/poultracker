package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sula.io/solevis/poultracker/internal/models"
	"git.sula.io/solevis/poultracker/internal/repositories"
	"github.com/go-chi/chi/v5"
)

func CreateCollectionHandler(w http.ResponseWriter, r *http.Request) {
	var collection models.Collection
	err := json.NewDecoder(r.Body).Decode(&collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	store := repositories.CollectionStore{}
	collection, err = store.Create(collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FindCollectionHandler(w http.ResponseWriter, r *http.Request) {
	laidDate := chi.URLParam(r, "laidDate")

	store := repositories.CollectionStore{}
	collection, err := store.FetchOne(laidDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FindAllCollectionsHandler(w http.ResponseWriter, _ *http.Request) {
	store := repositories.CollectionStore{}
	collections, err := store.FetchAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(collections)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateCollectionHandler(w http.ResponseWriter, r *http.Request) {
	var collection models.Collection
	err := json.NewDecoder(r.Body).Decode(&collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	store := repositories.CollectionStore{}
	affectedRows, err := store.Update(collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if affectedRows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteCollectionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	store := repositories.CollectionStore{}
	_, err = store.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
