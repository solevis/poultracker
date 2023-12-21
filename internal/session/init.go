package session

import (
	"log"
	"time"

	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager

func Init() *scs.SessionManager {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	log.Println("Loaded session manager")

	return sessionManager
}

func GetSessionManager() *scs.SessionManager {
	return sessionManager
}
