package routers

import (
	"github.com/webonise/csv_upload/app/controllers"
	"github.com/webonise/csv_upload/app/middlewares"
	"github.com/webonise/csv_upload/pkg/router"
)

// SrvRouter contains server HTTP Multiplexer
type SrvRouter struct {
	Router     router.Router
	Middleware middlewares.SrvAuthenticator
	Controller controllers.BaseController
}

// InitializeSrvRoutes intializes server routes
func (r *SrvRouter) InitializeSrvRoutes() {

	// REST APIs
	r.Router.Get("/api/ping", r.Middleware.Handle(r.Controller.Ping))

	// VIEWS
	r.Router.Get("/users", r.Middleware.RenderView(r.Controller.RenderUserView))
	r.Router.Post("/users", r.Middleware.RenderView(r.Controller.UserSignup))
	r.Router.Get("/upload", r.Middleware.RenderView(r.Controller.RenderUploadView))
	r.Router.Post("/upload", r.Middleware.Handle(r.Controller.UploadFile))
}
