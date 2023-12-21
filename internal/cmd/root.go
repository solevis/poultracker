package cmd

import (
	"fmt"
	"log"
	"net/http"

	"git.sula.io/solevis/poultracker/internal/config"
	"git.sula.io/solevis/poultracker/internal/database"
	"git.sula.io/solevis/poultracker/internal/migrations"
	"git.sula.io/solevis/poultracker/internal/router"
	"git.sula.io/solevis/poultracker/internal/session"
)

func Execute() int {
	var err error

	// load config from env
	err = config.Init()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	// connect to the database
	db, err := database.Init()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	defer db.Close()

	// init database schema
	err = migrations.Run()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	// setup session
	sessionManager := session.Init()

	// setup routes
	router := router.Init()

	// serve
	addr := fmt.Sprintf("%s:%d", config.Host(), config.Port())
	log.Printf("Running on http://%s\n", addr)
	err = http.ListenAndServe(addr, sessionManager.LoadAndSave(router))
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
