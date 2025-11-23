package graph

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// Keep wrappers to ensure compile if referenced
func MarshalTime(t time.Time) graphql.Marshaler { return graphql.MarshalTime(t) }
func UnmarshalTime(v interface{}) (time.Time, error) { return graphql.UnmarshalTime(v) }

