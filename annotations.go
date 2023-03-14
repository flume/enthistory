package enthistory

type UpdatedByKey string

func (UpdatedByKey) Name() string {
	return "UpdatedByKey"
}

type HistoryAnnotation string

func (HistoryAnnotation) Name() string {
	return "HistoryAnnotation"
}
