package main

import (
	"log"

	"fx.prodigy9.co/app"
	"fx.prodigy9.co/httpserver/controllers"
	"go.jonnrb.io/vanity"
)

func main() {
	handler := vanity.GitHubHandler(
		"go.hatari.fan",
		"hatari-electric", "go.hatari.fan", "https",
	)

	err := app.
		Build().
		DefaultAPIMiddlewares().
		Controllers(controllers.FromHandler("/*", handler)).
		Start()
	if err != nil {
		log.Fatalln(err)
	}
}
