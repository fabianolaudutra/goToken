package sync

import (
	"encoding/json"
	"net/http"
)

func responseJSON(res http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write([]byte(response))
}

func responseError(res http.ResponseWriter, code int, message string) {
	responseJSON(res, code, map[string]string{"error": message})
}
