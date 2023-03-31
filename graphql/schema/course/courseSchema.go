package course

import (
	"lms/enums"
	"lms/graphql/resolver"
	"lms/graphql/types"

	"github.com/graphql-go/graphql"
)

func GetCourses() *graphql.Field {
	return &graphql.Field{
		Type: types.CourseType,
		Args: graphql.FieldConfigArgument{
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"pageSize": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.GetCourses,
	}
}

func GetCourse() *graphql.Field {
	return &graphql.Field{
		Type: types.CourseType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolver.GetCourse,
	}
}

func CreateCourse() *graphql.Field {
	return &graphql.Field{
		Type: types.CourseType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"user_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"status": &graphql.ArgumentConfig{
				Type: enums.GetStatusEnumType(),
			},
		},
		Resolve: resolver.CreateCourse,
	}
}

func UpdateCourse() *graphql.Field {
	return &graphql.Field{
		Type: types.CourseType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"user_id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"status": &graphql.ArgumentConfig{
				Type: enums.GetStatusEnumType(),
			},
		},
		Resolve: resolver.UpdateCourse,
	}
}

func DeleteCourse() *graphql.Field {
	return &graphql.Field{
		Type: types.CourseType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: resolver.DeleteCourse,
	}
}
