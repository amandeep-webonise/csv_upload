package controllers

import (
	"github.com/webonise/csv_upload/app/services"
	"github.com/webonise/csv_upload/pkg/framework"
	"github.com/webonise/csv_upload/pkg/logger"
	"github.com/webonise/csv_upload/pkg/templates"
)

// Srv contains essential dependencies on which controllers rely
type Srv struct {
	Log       logger.Ilogger
	TplParser templates.ITemplateParser
	Service   services.ServiceProvider
}

// BaseController implements controller methods
type BaseController interface {
	Ping(w *framework.Response, r *framework.Request)
	RenderUserView(w *framework.Response, r *framework.Request)
	UserSignup(w *framework.Response, r *framework.Request)
	RenderUploadView(w *framework.Response, r *framework.Request)
	UploadFile(w *framework.Response, r *framework.Request)
}
