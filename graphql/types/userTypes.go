package types

import (
	"github.com/graphql-go/graphql"
)

type User struct {
    ID        int    `json:"id"`
    Phone     string `json:"phone"`
    Name      string `json:"name"`
    Password  string `json:"password"`
    Role      string `json:"role"`
    Status    string `json:"status"`
    CreatedAt string `json:"createdAt"`
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
        "phone": &graphql.Field{
            Type: graphql.String,
        },
        "password": &graphql.Field{
            Type: graphql.String,
        },
        "role": &graphql.Field{
            Type: graphql.String,
        },
        "status": &graphql.Field{
            Type: graphql.String,
        },
        "createdAt": &graphql.Field{
            Type: graphql.DateTime,
        },
    },
})

type Users []User
