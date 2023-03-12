package enthistory

import "database/sql/driver"

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
	for _, s := range opTypes {
		kinds = append(kinds, s)
	}
	return
}

func (op OpType) Value() (driver.Value, error) {
	return op.String(), nil
}

func (op OpType) String() string {
	return string(op)
}
