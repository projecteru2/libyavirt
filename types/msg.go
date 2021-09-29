package types

// Msg .
type Msg struct {
	Msg string
}

// NewMsg .
func NewMsg(msg string) Msg {
	return Msg{Msg: msg}
}

type WaitResult struct {
	Msg  string
	Code int64
}
