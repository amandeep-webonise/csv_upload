package middlewares

import (
	"net/http"

	"github.com/webonise/csv_upload/app/services"

	"github.com/webonise/csv_upload/pkg/framework"
	"github.com/webonise/csv_upload/pkg/monitoring"
)

// BaseMiddleware contains server middleware configurations
type BaseMiddleware struct {
	Service services.ServiceProvider
	Notify  monitoring.PanicNotifier
}

// SrvAuthenticator provides server middleware methods
type SrvAuthenticator interface {
	Handle(handler func(*framework.Response, *framework.Request)) http.HandlerFunc
	RenderView(viewHandler func(*framework.Response, *framework.Request)) http.HandlerFunc
}

// Handle will be serving only those requests that dont need to be authed
func (m *BaseMiddleware) Handle(handler func(*framework.Response, *framework.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer m.Notify.Capture(r)
		res := framework.NewResponse(w)
		req := framework.Request{Request: r}
		handler(&res, &req)
		res.Write()
	})
}

// RenderView renders a view
func (m *BaseMiddleware) RenderView(viewHandler func(*framework.Response, *framework.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer m.Notify.Capture(r)
		res := framework.NewResponse(w)
		req := framework.Request{Request: r}
		viewHandler(&res, &req)
	})
}
