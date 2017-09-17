package endpoints

import (
	"net/http"

	"github.com/hashknife/api/config"
)

type Frontender interface {
	http.Handler
}

// Frontend
type Frontend struct {
	path string
}

// NewFrontend
func NewFrontend(c *config.Config) Frontender {
	return &Frontend{
		path: *c.FrontendPath,
	}
}

// ServeHTTP is used to handle requests to the frontend
func (f *Frontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/html/index.html")
}
