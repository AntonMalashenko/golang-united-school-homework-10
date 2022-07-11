package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	var a, b int
	for k, v := range r.Header {
		switch strings.ToLower(k) {
		case "a":
			a, _ = strconv.Atoi(v[0])
		case "b":
			b, _ = strconv.Atoi(v[0])
		}
	}
	result := a + b
	w.Header().Set("a+b", strconv.Itoa(result))
}
