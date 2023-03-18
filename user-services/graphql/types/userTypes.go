package types

import (
	"github.com/graphql-go/graphql"
)

type User struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    CreatedAt string `json:"created_at"`
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
    Name: "User",
	Description: "User Type",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.Int,
        },
        "name": &graphql.Field{
            Type: graphql.String,
        },
        "email": &graphql.Field{
            Type: graphql.String,
        },
        "password": &graphql.Field{
            Type: graphql.String,
        },
        "created_at": &graphql.Field{
            Type: graphql.DateTime,
        },
    },
})

type Users []User

