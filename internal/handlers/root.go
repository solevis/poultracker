package handlers

import (
	"net/http"
	"time"

	"git.sula.io/solevis/poultracker/internal/models"
	"git.sula.io/solevis/poultracker/internal/repositories"
)

type HomeData struct {
	TodayCollection models.Collection
}

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	// get current collection of the day
	now := time.Now()
	store := repositories.CollectionStore{}
	todayCollection, _ := store.FetchOne(now.Format("2006-01-02"))
	data := HomeData{TodayCollection: todayCollection}

	// get & execute the template
	err := Template.ExecuteTemplate(w, "home", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func ChartHandler(w http.ResponseWriter, _ *http.Request) {
	// get & execute the template
	err := Template.ExecuteTemplate(w, "chart", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
