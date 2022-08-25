package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/Jonss/heroku-go-embed-template/pkg/config"
	"github.com/Jonss/heroku-go-embed-template/pkg/db"
	"github.com/Jonss/heroku-go-embed-template/pkg/server"
	"github.com/gorilla/mux"
)

//go:embed ui/build
var embeddedFiles embed.FS
var cfg config.Config


func main() {
	var err error
	cfg, err = config.LoadConfig(".")
	if err != nil {
		log.Fatalf("error getting config: error=(%v)", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}

	conn, err := db.NewConnection(cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate(conn, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}

	validator, err := server.NewValidator()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	srv := server.NewServer(
		router,
		cfg,
		validator,
	)
	srv.Routes()

	fmt.Printf("Starting Server... env: [%s].\n", cfg.Env)
	
	router.PathPrefix("/").Handler(http.FileServer(getFileSystem()))
	// http.Handle("/", )
	fmt.Println("Server started!")
	http.ListenAndServe(":"+port, router)

}

func getFileSystem() http.FileSystem {
	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(embeddedFiles, "ui/build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
