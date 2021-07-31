package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yadiiiig/weekend-challenge-one/internals/db"
	router "github.com/Yadiiiig/weekend-challenge-one/internals/router"
)

func main() {
	connection, err := db.ConnectDB("root:pass@(localhost:3306)/kma?parseTime=true")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	collection := router.InitRouter(connection)
	collection.InitRoutes()

	log.Fatal(http.ListenAndServe(":8080", collection.Router))
}
