package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	err_template "github.com/blattaria7/go-template/internal/errors"
)

type Handler struct {
	svc    Service
	logger *zap.Logger
}

func NewHandler(svc Service, logger *zap.Logger) *Handler {
	return &Handler{
		svc:    svc,
		logger: logger,
	}
}

type Service interface {
	Get(id string) (string, error)
}

type serviceResponse struct {
	Data  string                     `json:"data"`
	Error *err_template.ServiceError `json:"error,omitempty"`
}

// Healthcheck handler for checking the service life.
func (h *Handler) Healthcheck(w http.ResponseWriter, _ *http.Request) {
	var resp struct {
		Status string `json:"status,omitempty"`
	}

	resp.Status = "ok"

	data, err := json.Marshal(resp)
	if err != nil {
		h.logger.Error("error marshal data", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if _, err = w.Write(data); err != nil {
		h.logger.Error("error write data", zap.Error(err))
	}
}

func (h *Handler) Get(w http.ResponseWriter, req *http.Request) {
	var srvResponse serviceResponse

	params := mux.Vars(req)
	id := params["id"]

	result, err := h.svc.Get(id)
	switch {
	case err == nil:
	case errors.Is(err, err_template.ErrValueNotFound):
		srvResponse.Error = err_template.ErrValueNotFound
		srvResponse.Data = ""

		h.logger.Warn("id not found", zap.String("id", id))
		mustWriteErr(w, http.StatusNotFound, srvResponse)

		return
	default:
		srvResponse.Error = err_template.ErrInternalServerError
		srvResponse.Data = ""

		h.logger.Error("internal error", zap.Error(err))
		mustWriteErr(w, http.StatusInternalServerError, srvResponse)

		return
	}

	srvResponse.Data = result

	respBody, err := json.Marshal(srvResponse)
	if err != nil {
		srvResponse.Error = err_template.ErrInternalServerError
		srvResponse.Data = ""

		h.logger.Error("marshal data error", zap.Error(err))
		mustWriteErr(w, http.StatusInternalServerError, srvResponse)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(respBody)
	if err != nil {
		h.logger.Error("write response error", zap.Error(err))
	}
}

func mustWriteErr(w http.ResponseWriter, code int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	if _, err = w.Write(b); err != nil {
		panic(err)
	}
}
