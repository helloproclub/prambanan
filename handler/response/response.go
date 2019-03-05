package response

import (
	"encoding/json"
	"github.com/helloproclub/prambanan/errors"
	"net/http"
	"time"
)

const timeFormat = "2006-01-02T15:04:05"

// Success response
func Success(status int, w http.ResponseWriter, data interface{}) {
	meta := Meta{
		HTTPStatus: status,
	}
	SuccessMeta(status, w, data, meta)
}

// SuccessMeta success response with custom meta
func SuccessMeta(status int, w http.ResponseWriter, data interface{}, meta Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	loc := time.FixedZone("UTC+7", 7*60*60)
	meta.CurrentTime = time.Now().In(loc).Format(timeFormat)

	json.NewEncoder(w).Encode(&Body{
		Data: data,
		Meta: meta,
	})
}

// Message response with custom message
func Message(status int, w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	loc := time.FixedZone("UTC+7", 7*60*60)

	json.NewEncoder(w).Encode(&Body{
		Message: message,
		Meta: Meta{
			HTTPStatus:  status,
			CurrentTime: time.Now().In(loc).Format(timeFormat),
		},
	})
}

// Error response
func Error(w http.ResponseWriter, err error) {
	status := err.(*errors.Error).Status

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	loc := time.FixedZone("UTC+7", 7*60*60)

	json.NewEncoder(w).Encode(&Body{
		Error: err.(*errors.Error),
		Meta: Meta{
			HTTPStatus:  status,
			CurrentTime: time.Now().In(loc).Format(timeFormat),
		},
	})
}
