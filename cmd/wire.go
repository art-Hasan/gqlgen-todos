//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/art-Hasan/gqlgen-todos/graph"
)

func Service() (service, error) {
	wire.Build(
		newDB,
		newEnt,
		graph.NewResolver,
		newAddr,
		newRouter,
		newService,
	)
	return service{}, nil
}
