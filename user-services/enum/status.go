package enums

import "github.com/graphql-go/graphql"

var statusEnumType = graphql.NewEnum(graphql.EnumConfig{
	Name: "Status",
	Values: graphql.EnumValueConfigMap{
		"ACTIVE": &graphql.EnumValueConfig{
			Value: "ACTIVE",
		},
		"INACTIVE": &graphql.EnumValueConfig{
			Value: "INACTIVE",
		},
	},
	Description: "Status of the user",
})

func GetStatusEnumType() *graphql.Enum {
	return statusEnumType
}
