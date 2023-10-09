package files

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"

	"bjvanbemmel.nl/wopi/router"
	"bjvanbemmel.nl/wopi/storage"
	"github.com/go-chi/chi/v5"
)

//go:embed test.docx
var template string

var store storage.Storage = storage.New("./.storage")

func InitRoutes() {
	router.Router.Route("/wopi/files", func(r chi.Router) {
		r.Get("/{id}", getFileInfo)

		r.Get("/", listAllFiles)

		r.Post("/", createNewFile)

		r.Get("/{id}/contents", getFileContents)

		r.Post("/{id}/contents", putFileContents)

		r.Delete("/{id}", deleteFile)
	})
}

func getFileInfo(w http.ResponseWriter, r *http.Request) {
	file, content, err := store.Open(chi.URLParam(r, "id"), true)
	if err != nil {
		router.Json(w, err.Error())

		return
	}

	var response map[string]interface{} = make(map[string]interface{})
	response["BaseFileName"] = file.Name
	response["UserCanWrite"] = true
	response["Size"] = len(content)
    response["HideSaveOption"] = true

	router.Json(w, response)
}

func getFileContents(w http.ResponseWriter, r *http.Request) {
	_, raw, err := store.Open(chi.URLParam(r, "id"), true)
	if err != nil {
		router.Json(w, err.Error())

		return
	}
	fmt.Fprintln(w, string(raw))
}

func putFileContents(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	raw, _ := io.ReadAll(r.Body)

	var err error
	if store.Contains(chi.URLParam(r, "id")) {
		err = store.Update(chi.URLParam(r, "id"), raw)
	} else {
		_, err = store.Create(raw)
	}

	if err != nil {
		router.Json(w, err.Error())

		return
	}

    fmt.Println(len(raw))
	fmt.Fprintln(w, "200")
}

func createNewFile(w http.ResponseWriter, r *http.Request) {
	file, err := store.Create([]byte(template))
	if err != nil {
		router.Json(w, err.Error())

		return
	}

	router.Json(w, file)
}

func listAllFiles(w http.ResponseWriter, r *http.Request) {
	router.Json(w, store.Files)
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
	router.Json(w, store.Delete(chi.URLParam(r, "id")))
}
