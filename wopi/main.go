package main

import (
	"net/http"

	"bjvanbemmel.nl/wopi/files"
	"bjvanbemmel.nl/wopi/router"
)

func main() {
	router.New()
	files.InitRoutes()

	http.ListenAndServe(":80", router.Router)
}
