package handlers

import (
	"apimodels"
	"db"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func Search(rw http.ResponseWriter, r *http.Request) {
	var result2 map[string]map[string]interface{}
	var hits []apimodels.Message

	rw.Header().Set("Content-type", "application/json")

	query := chi.URLParam(r, "*")
	query = cleanQuery(query)

	result := db.Search(query)
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)

	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(body), &result2)

	jsonData, _ := json.Marshal(result2["hits"]["hits"])
	json.Unmarshal([]byte(jsonData), &hits)
	apimodels.SendData(rw, hits)

}

func cleanQuery(query string) string {
	clean := query

	clean = strings.ReplaceAll(clean, "&", " ")
	clean = strings.ReplaceAll(clean, "\"", "\\\"")

	return clean
}
