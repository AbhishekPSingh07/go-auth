package auth

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	key = "randomAssString"
	MaxAge = 86400 *30
	IsProd = false
)

func NewAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading env file")
	}

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleCLientSecret := os.Getenv("GOOGLE_CLIENT_SECRET_ID")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)
	
	store.Options.Path="/"
	store.Options.HttpOnly = true
	store.Options.Secure =IsProd

	gothic.Store = store

	
	goth.UseProviders(
		google.New(googleClientId, googleCLientSecret, "http://localhost:3000/auth/google/callback"),
	)
	
}