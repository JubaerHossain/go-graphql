package schema

import (
	"lms/graphql/resolver"
	"lms/graphql/types"

	"github.com/graphql-go/graphql"
)

func Login() *graphql.Field {
	return &graphql.Field{
		Type: types.AuthType,
		Args: graphql.FieldConfigArgument{
			"phone": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: resolver.Login,
	}
}

func RefreshToken() *graphql.Field {
	return &graphql.Field{
		Type: types.AuthType,
		Args: graphql.FieldConfigArgument{
			"refreshToken": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: resolver.RefreshToken,
	}
}

func Logout() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"refreshToken": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: resolver.Logout,
	}
}

func GetAuthUser() *graphql.Field {
	return &graphql.Field{
		Type:    types.UserType,
		Resolve: resolver.GetAuthUser,
	}
}
