package controllers

import (
	"gin-gonic-google-auth/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func GoogleInit()  {
	goth.UseProviders(google.New(config.GetEnvValue("GOOGLE_KEY"), config.GetEnvValue("GOOGLE_SECRET"), "http://localhost:3000/auth/google/callback"))
}