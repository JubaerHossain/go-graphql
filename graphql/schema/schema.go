package schema

import (
	"github.com/graphql-go/graphql"
	// Import the user and course schema files using package aliases.
	"lms/graphql/schema/course"
	"lms/graphql/schema/user"
)

// Define the root query object.
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	// Merge the user and course queries into the root query object.
	Fields: graphql.Fields{
		// Merge the user and course queries into the root query object.
		"users":   user.GetUsers(),     // Use the shorthand syntax to define the fields.
		"user":    user.GetUser(),      // Use the shorthand syntax to define the fields.
		"courses": course.GetCourses(), // Use the shorthand syntax to define the fields.
		"course":  course.GetCourse(),  // Use the shorthand syntax to define the fields.
	},
})

// Define the root mutation object.
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	// Merge the user and course mutations into the root mutation object.
	Fields: graphql.Fields{
		// Use the shorthand syntax to define the fields.
		"createUser":   user.CreateUser(),     // Use the shorthand syntax to define the fields.
		"updateUser":   user.UpdateUser(),     // Use the shorthand syntax to define the fields.
		"deleteUser":   user.DeleteUser(),     // Use the shorthand syntax to define the fields.
		"createCourse": course.CreateCourse(), // Use the shorthand syntax to define the fields.
		"updateCourse": course.UpdateCourse(), // Use the shorthand syntax to define the fields.
		"deleteCourse": course.DeleteCourse(), // Use the shorthand syntax to define the fields.

	},
})

// Create the schema from the root query and mutation objects.
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})
