package handler

import (
	"net/http"
)

func GetExampleRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("alex"))
}
