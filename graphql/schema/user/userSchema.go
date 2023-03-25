package user

import (
	"lms/enums"
	"lms/graphql/resolver"
	"lms/graphql/types"

	"github.com/graphql-go/graphql"
)

func GetUsers() *graphql.Field {
	return &graphql.Field{
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
	}
}

func GetUser() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.GetUser,
	}
}

func CreateUser() *graphql.Field {
	return &graphql.Field{
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
	}
}

func UpdateUser() *graphql.Field {
	return &graphql.Field{
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
	}
}

func DeleteUser() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: resolver.DeleteUser,
	}
}
