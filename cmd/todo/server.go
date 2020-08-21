package todo

import (
	"database/sql"

	"github.com/julienschmidt/httprouter"

	"github.com/vugu-examples/todo/cmd/todo/ctrl/todo"
	"github.com/vugu-examples/todo/cmd/todo/store"
)

type Server struct {
	DB     *sql.DB
	Router *httprouter.Router
	Store  *store.Store
	Ctrl   *todo.Ctrl
}

func NewServer(DB *sql.DB, router *httprouter.Router, store *store.Store, ctrl *todo.Ctrl) *Server {
	return &Server{DB: DB, Router: router, Store: store, Ctrl: ctrl}
}
