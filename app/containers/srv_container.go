package containers

import (
	"database/sql"

	"github.com/webonise/csv_upload/app/configs"
	"github.com/webonise/csv_upload/app/controllers"
	"github.com/webonise/csv_upload/app/middlewares"
	"github.com/webonise/csv_upload/app/models"
	"github.com/webonise/csv_upload/app/routers"
	"github.com/webonise/csv_upload/app/services"
	"github.com/webonise/csv_upload/pkg/logger"
	"github.com/webonise/csv_upload/pkg/monitoring"
	"github.com/webonise/csv_upload/pkg/router"
	"github.com/webonise/csv_upload/pkg/templates"
)

// Server contains server container
// Essential dependencies on which server application rely
type Server struct {
	Router    *router.Multiplexer
	Log       logger.Ilogger
	TplParser templates.ITemplateParser
	Cfg       *configs.ServerConfig
	DB        *sql.DB
}

// InitializeSrv initializes server application
func (s *Server) InitializeSrv() {

	serv := &services.Service{
		User:     &models.UserServiceImpl{DB: s.DB},
		Log:      s.Log,
		Employee: &models.EmployeeServiceImpl{DB: s.DB},
	}

	sr := &routers.SrvRouter{
		Router: s.Router,
		Middleware: &middlewares.BaseMiddleware{
			Service: serv,
			Notify: &monitoring.AirbrakeConfig{
				APIKey:      s.Cfg.AirbrakeAPIKey,
				Endpoint:    s.Cfg.AirbrakeServer,
				Environment: s.Cfg.Environment,
			},
		},
		Controller: &controllers.Srv{
			Log:       s.Log,
			TplParser: s.TplParser,
			Service:   serv,
		},
	}

	sr.InitializeSrvRoutes()
}
