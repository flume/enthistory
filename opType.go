package enthistory

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
)

type OpType string

const (
	OpTypeInsert OpType = "INSERT"
	OpTypeUpdate OpType = "UPDATE"
	OpTypeDelete OpType = "DELETE"
)

var opTypes = []string{
	OpTypeInsert.String(),
	OpTypeUpdate.String(),
	OpTypeDelete.String(),
}

// Values provides list valid values for Enum.
func (OpType) Values() (kinds []string) {
	kinds = append(kinds, opTypes...)
	return
}

func (op OpType) Value() (driver.Value, error) {
	return op.String(), nil
}

func (op OpType) String() string {
	return string(op)
}

func (op OpType) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.Quote(op.String())))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (op *OpType) UnmarshalGQL(val any) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*op = OpType(str)

	switch *op {
	case OpTypeInsert, OpTypeUpdate, OpTypeDelete:
		return nil
	default:
		return fmt.Errorf("%s is not a valid history operation type", str)
	}
}

func (op OpType) Name() (string, error) {
	switch op {
	case OpTypeInsert:
		return "OpTypeInsert", nil
	case OpTypeUpdate:
		return "OpTypeUpdate", nil
	case OpTypeDelete:
		return "OpTypeDelete", nil
	default:
		return "", fmt.Errorf("%s is not a valid history operation type", op)
	}
}
