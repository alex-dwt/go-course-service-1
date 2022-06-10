package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"alex/test/internal/http/request"
	"alex/test/internal/http/response"
)

func PostExampleRoute(w http.ResponseWriter, r *http.Request) {
	// parse request
	var body request.PostExampleRequest
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Decode error: "+err.Error())
		return
	}

	// send response
	result := response.PostExampleResponse{
		ID:  "id-id",
		Log: fmt.Sprintf("%+v", body),
	}

	encoded, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(encoded)
}
