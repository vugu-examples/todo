package todo

import (

	"github.com/julienschmidt/httprouter"
)

func NewRouter(ctrl *Controller) *httprouter.Router {

	r := httprouter.New()
	r.HandlerFunc("GET", "/", ctrl.List)
	return r

}

