package src

type State int

const (
	STATE_VOID State = iota
	STATE_BLACK
	STATE_WHITE
)

type BanVector int

const (
	BAN_UP BanVector = iota
	BAN_DOWN
	BAN_LEFT
	BAN_RIGHT
	BAN_VOID
)
