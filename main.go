package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
)

//go:embed react-app/build
var embeddedFiles embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}

	fmt.Println("Starting Server")
	http.HandleFunc("/api/teste", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá, go heroku go embed template!")
	})
	http.Handle("/", http.FileServer(getFileSystem()))
	http.ListenAndServe(":"+port, nil)
}

func getFileSystem() http.FileSystem {

	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(embeddedFiles, "react-app/build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
