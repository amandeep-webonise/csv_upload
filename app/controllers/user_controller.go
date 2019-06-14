package controllers

import (
	"net/http"

	"github.com/webonise/csv_upload/app/models"
	"github.com/webonise/csv_upload/pkg/framework"
)

// RenderUserView renders users details
func (s *Srv) RenderUserView(w *framework.Response, r *framework.Request) {
	us, err := s.Service.FetchAllUsers()
	if err != nil {
		s.Log.Error(`failed to fetch users`, err)
		http.Error(w.ResponseWriter, `No users found`, http.StatusFound)
		return
	}

	c := &struct {
		Users []*models.User
		Flash *framework.Flash
	}{
		us,
		r.GetFlash(w.ResponseWriter),
	}

	tmplList := []string{
		"./web/views/user.html",
		"./web/layouts/flash.html",
	}
	res, err := s.TplParser.ParseTemplate(tmplList, c)
	if err != nil {
		s.Log.Error(err)
		http.Error(w.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	w.RenderHTML(res)
}

// UserSignup signup for a user
func (s *Srv) UserSignup(w *framework.Response, r *framework.Request) {
	w.SetFlash("This is a simple flash message demo", "success")
	w.Redirect("/users", r.Request)
}
