package schema

import (
	enums "user-services/enum"
	"user-services/graphql/resolver"
	"user-services/graphql/types"

	"github.com/graphql-go/graphql"
)

var UsersQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Query",
	Description: "Root of all queries",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(types.UserType),
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"pageSize": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
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
	Name:        "Mutation",
	Description: "Root of all mutations",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: types.UserType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"phone": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"role": &graphql.ArgumentConfig{
					Type: enums.GetRoleEnumType(),
				},
				"status": &graphql.ArgumentConfig{
					Type: enums.GetStatusEnumType(),
				},
			},
			Resolve: resolver.CreateUser,
		},
		"updateUser": &graphql.Field{
			Type: types.UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"phone": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: resolver.UpdateUser,
		},
		"deleteUser": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: resolver.DeleteUser,
		},
	},
})

var UsersSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    UsersQuery,
	Mutation: UsersMutation,
})
