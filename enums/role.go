package enums

import "github.com/graphql-go/graphql"

var roleType = graphql.NewEnum(graphql.EnumConfig{
	Name: "Role",
	Values: graphql.EnumValueConfigMap{
		"ADMIN": &graphql.EnumValueConfig{
			Value: "ADMIN",
		},
		"USER": &graphql.EnumValueConfig{
			Value: "USER",
		},
		"INSTRUCTOR": &graphql.EnumValueConfig{
			Value: "INSTRUCTOR",
		},
	},
	Description: "Role of the user",
})

func GetRoleEnumType() *graphql.Enum {
	return roleType
}
