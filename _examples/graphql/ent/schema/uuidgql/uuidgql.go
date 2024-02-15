package uuidgql

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalUUID(u uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(u.String()))
	})
}

func UnmarshalUUID(v interface{}) (u uuid.UUID, err error) {
	s, ok := v.(string)
	if !ok {
		return u, fmt.Errorf("invalid type %T, expect string", v)
	}
	return uuid.Parse(s)
}
