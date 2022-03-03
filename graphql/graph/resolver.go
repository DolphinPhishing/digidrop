package graph

import (
	"context"
	"digidrop/ent"
	"digidrop/graphql/auth"
	"digidrop/graphql/graph/generated"
	"digidrop/graphql/graph/model"

	"github.com/99designs/gqlgen/graphql"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

const DownloadFolder = "downloads"
const CompromisedPath = "digidrop.compromised"

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	GQLConfig := generated.Config{
		Resolvers: &Resolver{
			client: client,
		},
	}

	GQLConfig.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []model.RoleLevel) (res interface{}, err error) {
		currentUser, err := auth.ForContext(ctx)

		if err != nil {
			return nil, nil
		}

		for _, role := range roles {
			if role.String() == currentUser.Type.String() {
				return next(ctx)
			}
		}

		return nil, nil

	}

	return generated.NewExecutableSchema(GQLConfig)
}
