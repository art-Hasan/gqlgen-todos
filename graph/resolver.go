package graph

import (
	"github.com/art-Hasan/gqlgen-todos/ent"
	"github.com/art-Hasan/gqlgen-todos/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ent *ent.Client
}

func NewResolver(ent *ent.Client) generated.ResolverRoot {
	return &Resolver{ent: ent}
}
