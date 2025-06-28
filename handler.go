package package_http

import "github.com/go-chi/chi/v5"

type Handler interface {
	Register(router chi.Router)
}
