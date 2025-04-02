package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/heismyke/backend_hospital/internal/app"
)

func SetupRoutes(myApp *app.Application) *chi.Mux{
	r := chi.NewRouter()
	r.Get("/health", myApp.CheckHealthStatus)	

	r.Post("/api/v1/register", myApp.UserHandler.HandleUserFunc)
	return r
}