package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"test-go/internal/handlers"
	"test-go/pkg/logging"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/users", h.GetList)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("list of users"))
}
