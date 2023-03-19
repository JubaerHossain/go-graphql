package schema

import (
	"course-services/enums"
	"course-services/graphql/resolver"
	"course-services/graphql/types"

	"github.com/graphql-go/graphql"
)

var CourseQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Course Query",
	Description: "Root of all queries",
	Fields: graphql.Fields{
		"courses": &graphql.Field{
			Type: types.CourseType,
			Args: graphql.FieldConfigArgument{
				"page":     &graphql.ArgumentConfig{Type: graphql.Int},
				"pageSize": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: resolver.GetCourses,
		},
		"course": &graphql.Field{
			Type: types.CourseType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolver.GetCourse,
		},
	},
})

var CourseMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Mutation",
	Description: "Root of all mutations",
	Fields: graphql.Fields{
		"createCourse": &graphql.Field{
			Type: types.CourseType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"status": &graphql.ArgumentConfig{
					Type: enums.GetStatusEnumType(),
				},
			},
			Resolve: resolver.CreateCourse,
		},
		"updateCourse": &graphql.Field{
			Type: types.CourseType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"status": &graphql.ArgumentConfig{
					Type: enums.GetStatusEnumType(),
				},
			},
			Resolve: resolver.UpdateCourse,
		},
		"deleteCourse": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: resolver.DeleteCourse,
		},
	},
})

var CourseSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    CourseQuery,
	Mutation: CourseMutation,
})
