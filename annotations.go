package enthistory

type UserIdKey string

func (UserIdKey) Name() string {
	return "UserIdKey"
}

type HistoryAnnotation string

func (HistoryAnnotation) Name() string {
	return "HistoryAnnotation"
}
