package todo

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/vugu-examples/todo/cmd/todo/todo_item_store"
)

type Server struct {
	DB * sql.DB
	Router *httprouter.Router
	Store *todo_item_store.ToDoItemStore
}

type Controller struct{}

func (ctrl *Controller) List(w http.ResponseWriter, r *http.Request){}