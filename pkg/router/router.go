package router

import (
	"net/http"

	"github.com/go-zoo/bone"
)

// Router encapsulates router methods
type Router interface {
	Get(path string, h http.Handler) *bone.Route
	Put(path string, h http.Handler) *bone.Route
	Post(path string, h http.Handler) *bone.Route
	Delete(path string, h http.Handler) *bone.Route
}

// Multiplexer contains router multiplexer
type Multiplexer struct {
	Mux *bone.Mux
}

// NewRouter creates and returns new router
func NewRouter() *bone.Mux {
	return bone.New()
}

// Get returns GET router method for use
func (m *Multiplexer) Get(path string, h http.Handler) *bone.Route {
	return m.Mux.Get(path, h)
}

// Put return bone put method for router
func (m *Multiplexer) Put(path string, h http.Handler) *bone.Route {
	return m.Mux.Put(path, h)
}

// Post return bone post method for router
func (m *Multiplexer) Post(path string, h http.Handler) *bone.Route {
	return m.Mux.Post(path, h)
}

// Delete return bone delete method for router
func (m *Multiplexer) Delete(path string, h http.Handler) *bone.Route {
	return m.Mux.Delete(path, h)
}
