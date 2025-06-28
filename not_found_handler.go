package package_http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	slog "github.com/salamsites/package-log"
)

type handler struct {
	ctx    context.Context
	logger *slog.Logger
}

// NotFoundHandler returns a new handler instance
func NotFoundHandler(ctx context.Context, logger *slog.Logger) *handler {
	return &handler{
		ctx:    ctx,
		logger: logger,
	}
}

// Register sets the custom NotFound handler for the chi router
func (h *handler) Register(router *chi.Mux) {
	router.NotFound(h.NotFound)
}

// NotFound is the handler for unknown routes
func (h *handler) NotFound(w http.ResponseWriter, r *http.Request) {
	h.logger.Error("middleware not found")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 not found"))
}
