//+build wireinject

package todo

import (
	"github.com/google/wire"
)

func Setup() (*App, error) {

	wire.Build(
		NewController,
		NewDBConnStr,
		NewDBDriverName,
		NewRouter,
		NewToDoApp,
		//todo_item_store.NewToDoItemStore,
	)

	return &App{}, nil
}
