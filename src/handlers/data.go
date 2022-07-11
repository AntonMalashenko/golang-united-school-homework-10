package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func DataHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	bodyString := string(body)
	w.Write([]byte(fmt.Sprintf("I got message:\n%s", bodyString)))
}
