// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package todo

import (
	"github.com/vugu-examples/todo/cmd/todo/ctrl/todo"
	"github.com/vugu-examples/todo/cmd/todo/store"
)

// Injectors from wire.go:

func Setup() (*App, error) {
	dbConnStr := NewDBConnStr()
	dbDriverName := NewDBDriverName()
	db := NewDBConn(dbConnStr, dbDriverName)
	toDoItemStore := store.NewToDoItemStore(db)
	ctrl := todo.NewCtrlToDo(toDoItemStore)
	router := todo.NewRouter(ctrl)
	app, err := NewToDoApp(dbConnStr, dbDriverName, router)
	if err != nil {
		return nil, err
	}
	return app, nil
}
