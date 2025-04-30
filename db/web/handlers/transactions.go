package handlers

import (
	"fmt"
	"net/http"
)

func Transtestfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint /transactions")
}
