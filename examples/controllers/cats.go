package controllers

import (
	"github.com/joncalhoun/viewcon"
	"github.com/joncalhoun/viewcon/examples/views"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CatsController struct {
	viewcon.Controller
}

var Cats CatsController

func (self CatsController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {

	// Pretend to lookup cats in some way.
	cats := []string{"heathcliff", "garfield"}

	// render the view
	return views.Cats.Index.Render(w, cats)
}

func (self CatsController) Feed(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	// render the view
	return views.Cats.Feed.Render(w, nil)
}

func (self *CatsController) ReqKey(a viewcon.Action) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.FormValue("key") != "topsecret" {
			http.Error(w, "Invalid key.", http.StatusUnauthorized)
		} else {
			self.Controller.Perform(a)(w, r, ps)
		}
	})
}
