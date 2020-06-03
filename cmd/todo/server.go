package todo

import (
	"database/sql"
	"encoding/json"
	"io"
	"mime"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/vugu-examples/todo/cmd/todo/todo_item_store"
)

type Server struct {
	DB     *sql.DB
	Router *httprouter.Router
	Store  *todo_item_store.ToDoItemStore
}

type Controller struct {
	Store *todo_item_store.ToDoItemStore
}

func NewController(toDoStore *todo_item_store.ToDoItemStore) *Controller {
	return &Controller{
		Store: toDoStore,
	}
}

func (ctrl *Controller) List(w http.ResponseWriter, r *http.Request) {

}

func (ctrl *Controller) GetOne(w http.ResponseWriter, r *http.Request) {}

func (ctrl *Controller) Create(w http.ResponseWriter, r *http.Request) {

	contentType, _, err := mime.ParseMediaType(r.Header.Get("content-type"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if contentType != "application/json" {
		http.Error(w, "unsupported media type", http.StatusUnsupportedMediaType) //need to figure out better error handling than this
		return
	}

	var obj todo_item_store.ToDoItem

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		http.Error(w, "request body must only contain a single JSON object", http.StatusBadRequest)
		return
	}

	_, err = ctrl.Store.Insert().Object(&obj).ExecContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(&obj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	return

}

func (ctrl *Controller) Delete(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	ToDoID := params.ByName("id")

	if len(ToDoID) == 0 {
		http.Error(w, "missing id param", 500)
		return
	}

	var err error

	_, err = ctrl.Store.Delete().Record(&todo_item_store.ToDoItem{
		ID: ToDoID,
	}).ExecContext(r.Context())

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	return

}
