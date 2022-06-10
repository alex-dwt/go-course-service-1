package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"alex/test/internal/http/response"
)

func GetExampleRoute(w http.ResponseWriter, r *http.Request) {
	result := response.GenerateFakeGetResponse()

	encoded, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(encoded)
}
