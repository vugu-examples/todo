//+build wireinject

package todo

import "github.com/google/wire"

func Wire() (*ToDoApp, error) {

	wire.Build(

		NewRouter,
		)
}