package todo

import (

	"github.com/julienschmidt/httprouter"
)

func NewRouter(ctrl *Controller) *httprouter.Router {

	r := httprouter.New()
	r.HandlerFunc("GET", "/", ctrl.List)
	r.HandlerFunc("GET", "/todo/:id", ctrl.GetOne)
	r.HandlerFunc("POST", "/todo", ctrl.Create)
	r.HandlerFunc("DELETE", "/todo/:id", ctrl.Delete)
	return r

}

