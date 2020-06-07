//+build wireinject

package todo

import (
	"github.com/google/wire"

	"github.com/vugu-examples/todo/cmd/todo/ctrl/todo"
	"github.com/vugu-examples/todo/cmd/todo/store"
)

func Setup() (*App, error) {

	wire.Build(
		todo.NewCtrlToDo,
		NewDBConn,
		NewDBConnStr,
		NewDBDriverName,
		todo.NewRouter,
		NewToDoApp,
		store.NewToDoItemStore,
	)

	return &App{}, nil
}
