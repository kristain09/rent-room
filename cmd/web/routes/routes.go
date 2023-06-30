package routes

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/kristain09/rent-room/pkg/config"
	"github.com/kristain09/rent-room/pkg/handlers"
)

func Router(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
