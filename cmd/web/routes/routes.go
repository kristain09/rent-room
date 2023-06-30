package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kristain09/rent-room/pkg/config"
	"github.com/kristain09/rent-room/pkg/handlers"
)

func Router(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
