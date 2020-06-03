//+build wireinject

package todo

import (
	"github.com/google/wire"

	"github.com/vugu-examples/todo/cmd/todo/todo_item_store"
)

func Setup() (*App, error) {

	wire.Build(
		NewController,
		NewDBConn,
		NewDBConnStr,
		NewDBDriverName,
		NewRouter,
		NewToDoApp,
		todo_item_store.NewToDoItemStore,
	)

	return &App{}, nil
}
