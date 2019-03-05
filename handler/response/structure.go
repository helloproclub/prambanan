package response

import (
	"github.com/helloproclub/prambanan/errors"
)

// Meta store meta info
type Meta struct {
	HTTPStatus  int    `json:"http_status"`
	Limit       int    `json:"limit,omitempty"`
	Offset      int    `json:"offset,omitempty"`
	Total       int    `json:"total,omitempty"`
	CurrentTime string `json:"current_time,omitempty"`
}

// Body store data, errors, and meta
type Body struct {
	Data    interface{}   `json:"data,omitempty"`
	Error   *errors.Error `json:"error,omitempty"`
	Message string        `json:"message,omitempty"`
	Meta    Meta          `json:"meta"`
}
