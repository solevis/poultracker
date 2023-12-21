package router

import (
	"log"
	"net/http"
	"strings"

	"git.sula.io/solevis/poultracker/internal/handlers"
	"git.sula.io/solevis/poultracker/internal/session"
	"git.sula.io/solevis/poultracker/web/static"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var router *chi.Mux

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// don't block login page and static files
		if strings.HasPrefix(path, "/auth/") || strings.HasPrefix(path, "/static/") {
			next.ServeHTTP(w, r)
			return
		}

		sessionManager := session.GetSessionManager()
		authenticated := sessionManager.GetBool(r.Context(), session.KeyAuthenticated)

		if !authenticated {
			// return 401 Unauthorized for api requests
			if strings.HasPrefix(path, "/api/") {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// avoid small default text in body
			w.Header().Set("Content-Type", "")

			http.Redirect(w, r, "/auth/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Init() *chi.Mux {
	router = chi.NewRouter()

	// setup middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(authenticationMiddleware)

	// static route
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(static.Files()))))

	// setup routes
	router.Get("/auth/login", handlers.LoginHandler)
	router.Get("/auth/logout", handlers.LogoutHandler)
	router.Post("/auth/login", handlers.ValidateLoginHandler)

	router.Get("/", handlers.HomeHandler)
	router.Get("/chart", handlers.ChartHandler)

	router.Get("/api/collections", handlers.FindAllCollectionsHandler)
	router.Post("/api/collections", handlers.CreateCollectionHandler)
	router.Put("/api/collections", handlers.UpdateCollectionHandler)
	router.Get("/api/collections/{laidDate}", handlers.FindCollectionHandler)
	router.Delete("/api/collections/{id}", handlers.DeleteCollectionHandler)

	log.Println("Loaded routes")

	return router
}
