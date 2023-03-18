package schema

import (
	"user-services/graphql/types"
	"user-services/graphql/resolver"
	"github.com/graphql-go/graphql"
)

var UsersQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(types.UserType),
			Resolve: resolver.GetUsers,
		},
		"user": &graphql.Field{
			Type: types.UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolver.GetUser,
		},
	},
})

var UsersMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: types.UserType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: resolver.CreateUser,
		},
	},
})

var UsersSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    UsersQuery,
	Mutation: UsersMutation,
})



