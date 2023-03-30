package resolvers

import (
	"go-web/ent"
	generated "go-web/graph/generated"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

const maxQueryComplexity = 300

// NewConfig
func NewConfig(client *ent.Client) *generated.Config {
	return &generated.Config{
		Resolvers: &Resolver{
			client: client,
		},
	}
}

// NewGraphqlHandler
func NewGraphqlHandler(c *generated.Config, client *ent.Client) *handler.Server {
	if c == nil {
		panic("config is nil")
	}

	h := handler.NewDefaultServer(generated.NewExecutableSchema(*c))
	h.Use(entgql.Transactioner{TxOpener: client})
	h.Use(extension.FixedComplexityLimit(maxQueryComplexity))

	return h
}
