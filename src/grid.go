package src

type Grid struct {
	x     int
	y     int
	state State
}

func NewGrid(mx int, my int) *Grid {
	return &Grid{x: mx, y: my, state: STATE_VOID}
}

func (g *Grid) SetState(s State) {
	g.state = s
}

func (g *Grid) ClearState() {
	g.state = STATE_VOID
}

func (g *Grid) GetState() State {
	return g.state
}

func (g *Grid) EqualState(s State) bool {
	return g.state == s
}
