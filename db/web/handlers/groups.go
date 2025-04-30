package handlers

import (
	"fmt"
	"net/http"
)

func Groupstestfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint /groups")
}
