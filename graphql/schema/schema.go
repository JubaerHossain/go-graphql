package schema

import (
	"github.com/graphql-go/graphql"
)

// Define the root query object.
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	// Merge the user and course queries into the root query object.
	Fields: graphql.Fields{
		// Merge the user and course queries into the root query object.
		"users":   GetUsers(),   // Use the shorthand syntax to define the fields.
		"user":    GetUser(),    // Use the shorthand syntax to define the fields.
		"courses": GetCourses(), // Use the shorthand syntax to define the fields.
		"course":  GetCourse(),  // Use the shorthand syntax to define the fields.
	},
})

// Define the root mutation object.
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	// Merge the user and course mutations into the root mutation object.
	Fields: graphql.Fields{
		// Use the shorthand syntax to define the fields.
		"createUser":   CreateUser(),   // Use the shorthand syntax to define the fields.
		"updateUser":   UpdateUser(),   // Use the shorthand syntax to define the fields.
		"deleteUser":   DeleteUser(),   // Use the shorthand syntax to define the fields.
		"createCourse": CreateCourse(), // Use the shorthand syntax to define the fields.
		"updateCourse": UpdateCourse(), // Use the shorthand syntax to define the fields.
		"deleteCourse": DeleteCourse(), // Use the shorthand syntax to define the fields.

	},
})

// Create the schema from the root query and mutation objects.
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})
