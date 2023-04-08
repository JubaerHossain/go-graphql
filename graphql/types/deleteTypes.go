package types

import (
	"github.com/graphql-go/graphql"
)

var DeleteType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Delete",
	Description: "Delete Type",
	Fields: graphql.Fields{

		"status": &graphql.Field{
			Type: graphql.Boolean,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
