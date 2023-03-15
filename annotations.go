package enthistory

const (
	ValueTypeInt ValueType = iota
	ValueTypeString
)

type ValueType uint

func (ValueType) ValueType() string {
	return "ValueType"
}
