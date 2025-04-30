package handlers

import (
	"fmt"
	"net/http"
)

func Goalstestfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint /goals")
}
