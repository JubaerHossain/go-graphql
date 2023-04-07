package types

import (
	"errors"
	"lms/database"
	"lms/model"
	"reflect"

	"github.com/JubaerHossain/gosql"
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
			Type: UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				x := p.Source.(model.Course)
				p.Args["id"] = x.User
				user, err := gosql.FindByID(reflect.TypeOf(model.User{}), "users", p, database.DB)
				if err != nil {
					return nil, errors.New("no data found")
				}
				return user, nil
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
