package types

import "github.com/graphql-go/graphql"

var AuthType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Auth",
	Description: "Auth Type",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"refreshToken": &graphql.Field{
			Type: graphql.String,
		},
	},
})

