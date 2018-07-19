package model

type Result struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OK(msg string, data interface{}) Result {
	return Result{msg, data}
}

func NoData() Result {
	return Result{"no data", nil}
}

func ParameterError() Result {
	return Result{"parameter error", nil}
}
