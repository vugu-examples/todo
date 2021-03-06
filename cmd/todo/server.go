package todo

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/vugu-examples/todo/cmd/todo/todo_item_store"
)

type Server struct {
	DB     *sql.DB
	Router *httprouter.Router
	Store  *todo_item_store.ToDoItemStore
}

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) List(w http.ResponseWriter, r *http.Request) {}

func (ctrl *Controller) GetOne(w http.ResponseWriter, r *http.Request) {}

func (ctrl *Controller) Create(w http.ResponseWriter, r *http.Request) {}

func (ctrl *Controller) Delete(w http.ResponseWriter, r *http.Request) {}
