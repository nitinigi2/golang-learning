package controller

import (
	"encoding/json"
	"net/http"
)

func RequestBodyToObject(w http.ResponseWriter, r *http.Request, any interface{}) error {
	err := json.NewDecoder(r.Body).Decode(any)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // request could not be correctly parsed  (including the request entity/body)
	}
	return err
}
