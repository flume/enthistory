package enthistory

const (
	ValueTypeInt ValueType = iota
	ValueTypeString
)

type ValueType uint

func (ValueType) ValueType() string {
	return "ValueType"
}

type Annotations struct {
	Exclude   bool `json:"exclude,omitempty"`   // Will exclude history tracking for this schema
	IsHistory bool `json:"isHistory,omitempty"` // DO NOT APPLY TO ANYTHING EXCEPT HISTORY SCHEMAS
}

func (Annotations) Name() string {
	return "History"
}
