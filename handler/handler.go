package handler

import (
	"github.com/helloproclub/prambanan"
	"github.com/helloproclub/prambanan/handler/response"
	"net/http"
)

type Handler struct {
	Prambanan *prambanan.Prambanan
}

func NewHandler(prambanan *prambanan.Prambanan) (handler *Handler, err error) {
	handler = &Handler{
		Prambanan: prambanan,
	}

	return
}

// Ping check application status
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	response.Message(200, w, "Pong")
}

// DecodeNik handle decoding NIK given a valid NIK
func (h *Handler) DecodeNik(w http.ResponseWriter, r *http.Request) {
	nik := getQuery(r, "nik")
	result, err := h.Prambanan.DecodeNIK(r.Context(), nik)
	if err != nil {
		response.Error(w, err)
	}
	response.Success(200, w, result)
}

func getQuery(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
