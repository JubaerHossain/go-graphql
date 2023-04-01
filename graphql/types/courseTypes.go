package types

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var CourseType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Course",
	Description: "Course Type",
	Fields: graphql.Fields{

		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: graphql.NewList(UserType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				fmt.Println("user")

				fmt.Println(p.Source)

				sMap := p.Source.(map[string]interface{})
				id := sMap["user"]
				fmt.Println(id)

				// user, err := query.FindModel(reflect.TypeOf(model.User{}), "users", params)
				// if err != nil {
				// 	fmt.Println(err)
				// 	return nil, errors.New("no data found")
				// }

				// return user, nil

				return nil, nil

			},
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})
