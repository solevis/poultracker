package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"git.sula.io/solevis/poultracker/internal/config"
	"git.sula.io/solevis/poultracker/internal/session"
)

type LoginData struct {
	Success bool
	Csrf    string
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func isCredentialsValid(username string, password string) bool {
	return username == config.Username() && password == config.Password()
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	newToken, err := generateToken()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	sessionManager := session.GetSessionManager()
	sessionManager.Put(r.Context(), session.KeyCsrf, newToken)

	loginData := LoginData{
		Success: true,
		Csrf:    newToken,
	}

	err = Template.ExecuteTemplate(w, "login", loginData)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// revoke users authentication
	sessionManager := session.GetSessionManager()
	sessionManager.Put(r.Context(), session.KeyAuthenticated, false)

	// redirect to login page
	http.Redirect(w, r, "/auth/login", http.StatusFound)
}

func ValidateLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	csrf := r.FormValue("csrf")

	// Retrieve session stored CSRF
	sessionManager := session.GetSessionManager()
	sessionCsrf := sessionManager.GetString(r.Context(), session.KeyCsrf)

	if csrf == sessionCsrf && isCredentialsValid(username, password) {
		// First renew the session token...
		err := sessionManager.RenewToken(r.Context())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Set user as authenticated
		sessionManager.Put(r.Context(), session.KeyAuthenticated, true)

		// Remove csrf
		sessionManager.Remove(r.Context(), session.KeyCsrf)

		// Redirect to home
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Generate new CSRF token
	newToken, err := generateToken()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	sessionManager.Put(r.Context(), session.KeyCsrf, newToken)

	loginData := LoginData{
		Success: false,
		Csrf:    newToken,
	}

	err = Template.ExecuteTemplate(w, "login", loginData)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
