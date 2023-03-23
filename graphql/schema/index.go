package schema

import (
	"github.com/graphql-go/graphql"
)

var combinedSchema = graphql.MergeSchemas(graphql.MergeSchemasConfig{
	Schemas: []graphql.Schema{CourseSchema},
})
