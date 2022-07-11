package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func NameHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := strings.Split(r.RequestURI, "/")
	name := pathParams[len(pathParams)-1]
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}
