package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	// decoding json in data type
	log.Print(r.Body)
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// helper method for response
func (app *Config) writeJSON(w http.ResponseWriter, status int, data jsonResponse, headers ...http.Header) error {
	// converting jsonResponse strut to json
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// added logic for write response header
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// response body data
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

// helper method for error response
func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	payload := jsonResponse{
		Status:  false,
		Message: err.Error(),
	}
	return app.writeJSON(w, statusCode, payload)
}
