package handlers

import (
	"fmt"
	"net/http"
)

func PostTest(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Post test success")
}
