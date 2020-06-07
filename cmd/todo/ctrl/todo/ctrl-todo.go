package todo

import (
	"encoding/json"
	"io"
	"mime"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/vugu-examples/todo/cmd/todo/store"
)

type Ctrl struct {
	Store *store.ToDoItemStore
}

func NewCtrlToDo(toDoStore *store.ToDoItemStore) *Ctrl {
	return &Ctrl{
		Store: toDoStore,
	}
}

func NewRouter(ctrl *Ctrl) *httprouter.Router {

	r := httprouter.New()
	r.HandlerFunc("GET", "/", ctrl.List)
	r.HandlerFunc("GET", "/todo/:id", ctrl.GetOne)
	r.HandlerFunc("POST", "/todo", ctrl.Create)
	r.HandlerFunc("DELETE", "/todo/:id", ctrl.Delete)
	return r

}

func (ctrl *Ctrl) List(w http.ResponseWriter, r *http.Request) {

}

func (ctrl *Ctrl) GetOne(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id := params.ByName("id")

	obj, err := ctrl.Store.Select().Where("id", id).GetOne()
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&obj); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	return

}

func (ctrl *Ctrl) Create(w http.ResponseWriter, r *http.Request) {

	contentType, _, err := mime.ParseMediaType(r.Header.Get("content-type"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if contentType != "application/json" {
		http.Error(w, "unsupported media type", http.StatusUnsupportedMediaType) //need to figure out better error handling than this
		return
	}

	var obj store.ToDoItem

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

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(&obj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	return

}

func (ctrl *Ctrl) Delete(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	ToDoID := params.ByName("id")

	if len(ToDoID) == 0 {
		http.Error(w, "missing id param", 500)
		return
	}

	var err error

	_, err = ctrl.Store.Delete().Record(&store.ToDoItem{
		ID: ToDoID,
	}).ExecContext(r.Context())

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	return

}
