package types

type EventMessage struct {
	ID       string
	Type     string
	Action   string
	TimeNano int64
}
