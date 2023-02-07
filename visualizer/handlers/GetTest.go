package handlers

import (
	"fmt"
	"net/http"
)

func GetTest(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Get test success")
}
