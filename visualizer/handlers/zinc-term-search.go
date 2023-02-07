package handlers

import (
	"db"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func Search(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")

	query := chi.URLParam(r, "*")
	query = cleanQuery(query)

	result := db.Search(query)
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(rw, string(body))

}

func cleanQuery(query string) string {
	clean := query

	clean = strings.ReplaceAll(clean, "&", " ")
	clean = strings.ReplaceAll(clean, "\"", "\\\"")

	return clean
}
