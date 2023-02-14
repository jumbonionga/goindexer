package main

import (
	"fmt"
	"handlers"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	port := 3000

	router.Get("/api/example/", handlers.GetTest)
	router.Post("/api/example/", handlers.PostTest)
	router.Get("/api/search/*", handlers.Search)

	fmt.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), router))
}
