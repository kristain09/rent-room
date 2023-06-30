package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kristain09/rent-room/pkg/config"
	"github.com/kristain09/rent-room/pkg/handlers"
	"github.com/kristain09/rent-room/pkg/render"
)

const (
	PORTNUMBER = ":8080"
)

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("cannot create a template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORTNUMBER))
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
