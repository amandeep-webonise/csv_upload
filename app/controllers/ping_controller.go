package controllers

import "github.com/webonise/csv_upload/pkg/framework"

// Ping returns pong response
func (s *Srv) Ping(w *framework.Response, r *framework.Request) {
	s.Log.Info("Hello From Log Side")
	w.Message("PONG")
}
