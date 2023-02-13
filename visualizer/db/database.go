package db

import (
	"log"
	"net/http"
	"strings"
)

const url = "http://localhost:4080/api/enron_test"
const user = "admin"
const password = "Complexpass#123"

func Search(querystring string) *http.Response {
	var response *http.Response
	var source string
	if strings.Contains(querystring, "_id") {
		source = ""
	} else {
		source = `"_id","Date","From","To","Subject"`
	}
	query := `{
			"search_type": "querystring",
			"query": {
			"term": "` + querystring + `"
			},
			"sort_fields":["-@timestamp"],
			"_source": [` + source + `]
	}`

	if req, err := http.NewRequest("POST", url+"/_search", strings.NewReader(query)); err != nil {
		log.Fatal(err)
	} else {
		req.SetBasicAuth(user, password)
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("User-Agent", "Go visualize go!")

		if resp, err := http.DefaultClient.Do(req); err != nil {
			log.Fatal(err)
		} else {
			response = resp
		}
	}
	return response
}
