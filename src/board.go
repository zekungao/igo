package src

import (
	"bytes"
	"fmt"
)

const _length int = 19
const _width int = 19

// type State int

// const (
// 	STATE_VOID State = iota
// 	STATE_BLACK
// 	STATE_WHITE
// )

// type BanVector int

// const (
// 	BAN_UP BanVector = iota
// 	BAN_DOWN
// 	BAN_LEFT
// 	BAN_RIGHT
// 	BAN_VOID
// )

// type Grid struct {
// 	x     int
// 	y     int
// 	state State
// }

// func NewGrid(mx int, my int) *Grid {
// 	return &Grid{x: mx, y: my, state: STATE_VOID}
// }

// func (b *Grid) SetState(s State) {
// 	b.state = s
// }

// func (b *Grid) ClearState() {
// 	b.state = STATE_VOID
// }

// func (b *Grid) GetState() State {
// 	return b.state
// }

// func (b *Grid) EqualState(s State) bool {
// 	if b.state == s {
// 		return true
// 	}
// 	return false
// }

type Board struct {
	originX int
	originY int
	Length  int
	Width   int
	board   [_length][_width]Grid
}

func NewBoard() *Board {
	chessBoard := [_length][_width]Grid{}
	for i := 0; i < _width; i++ {
		for j := 0; j < _length; j++ {
			chessBoard[i][j] = *NewGrid(i, j)
		}
	}
	return &Board{originX: -1, originY: -1, Length: _length, Width: _width, board: chessBoard}
}

func (b *Board) SetGO(x int, y int, state State) {
	// fmt.Println(b.board[x][y].state)
	b.board[x][y].SetState(state)
	// fmt.Println(b.board[x][y].state)
}

func (b *Board) JudgeDA(x int, y int, originState State, ban BanVector) bool {
	if x == b.originX && y == b.originY {
		return false
	}
	if b.board[x][y].EqualState(STATE_VOID) {
		return true
	}
	if !b.board[x][y].EqualState(originState) {
		return false
	}
	return b.DeathAlive(x, y, originState, ban)
}

func (b *Board) DeathAlive(x int, y int, originState State, ban BanVector) bool {
	//true:alive, false:death
	if x != 0 && ban != BAN_LEFT {
		if b.JudgeDA(x-1, y, originState, BAN_RIGHT) {
			return true
		}
	}
	if x != _width-1 && ban != BAN_RIGHT {
		if b.JudgeDA(x+1, y, originState, BAN_LEFT) {
			return true
		}
	}
	if y != 0 && ban != BAN_UP {
		if b.JudgeDA(x, y-1, originState, BAN_DOWN) {
			return true
		}
	}
	if y != _length-1 && ban != BAN_DOWN {
		if b.JudgeDA(x, y+1, originState, BAN_UP) {
			return true
		}
	}
	return false
}

func (b *Board) ResetOrigin() {
	b.originX = -1
	b.originY = -1
}

func (b *Board) CanEat(x int, y int, originState State) bool {
	b.originX = x
	b.originY = y
	if x != 0 && !b.board[x-1][y].EqualState(originState) {
		if !b.DeathAlive(x-1, y, b.board[x-1][y].GetState(), BAN_RIGHT) {
			b.ResetOrigin()
			return true
		}
	}
	if x != _width-1 && !b.board[x+1][y].EqualState(originState) {
		if !b.DeathAlive(x+1, y, b.board[x+1][y].GetState(), BAN_LEFT) {
			b.ResetOrigin()
			return true
		}
	}
	if y != 0 && !b.board[x][y-1].EqualState(originState) {
		if !b.DeathAlive(x, y-1, b.board[x][y-1].GetState(), BAN_DOWN) {
			b.ResetOrigin()
			return true
		}
	}
	if y != _length-1 && !b.board[x][y+1].EqualState(originState) {
		if !b.DeathAlive(x, y+1, b.board[x][y+1].GetState(), BAN_UP) {
			b.ResetOrigin()
			return true
		}
	}
	b.ResetOrigin()
	return false
}

func (b *Board) CanDrop(x int, y int, originState State) bool {
	if b.CanEat(x, y, originState) {
		return true
	}
	if b.DeathAlive(x, y, originState, BAN_VOID) {
		return true
	}
	return false
}

func (b *Board) Eating(x int, y int, ban BanVector) {
	originState := b.board[x][y].GetState()

	if x != 0 && b.board[x-1][y].EqualState(originState) && ban != BAN_LEFT {
		b.Eating(x-1, y, BAN_RIGHT)
	}
	if x != _width-1 && b.board[x+1][y].EqualState(originState) && ban != BAN_RIGHT {
		b.Eating(x+1, y, BAN_LEFT)
	}
	if y != 0 && b.board[x][y-1].EqualState(originState) && ban != BAN_DOWN {
		b.Eating(x, y-1, BAN_UP)
	}
	if y != _length-1 && b.board[x][y+1].EqualState(originState) && ban != BAN_UP {
		b.Eating(x, y+1, BAN_DOWN)
	}
	b.board[x][y].ClearState()
}

func (b *Board) Eat(x int, y int) {
	originState := b.board[x][y].GetState()
	if !b.CanEat(x, y, originState) {
		fmt.Println("not eat")
		return
	}
	if x != 0 && !b.board[x-1][y].EqualState(originState) {
		if !b.DeathAlive(x-1, y, b.board[x-1][y].GetState(), BAN_RIGHT) {
			b.Eating(x-1, y, BAN_RIGHT)
		}
	}
	if x != _width-1 && !b.board[x+1][y].EqualState(originState) {
		if !b.DeathAlive(x+1, y, b.board[x+1][y].GetState(), BAN_LEFT) {
			b.Eating(x+1, y, BAN_LEFT)
		}
	}
	if y != 0 && !b.board[x][y-1].EqualState(originState) {
		if !b.DeathAlive(x, y-1, b.board[x][y-1].GetState(), BAN_DOWN) {
			b.Eating(x, y-1, BAN_UP)
		}
	}
	if y != _length-1 && !b.board[x][y+1].EqualState(originState) {
		if !b.DeathAlive(x, y+1, b.board[x][y+1].GetState(), BAN_UP) {
			b.Eating(x, y+1, BAN_DOWN)
		}
	}
}

func (b *Board) Drop(x int, y int, state State) bool {
	if !b.CanDrop(x, y, state) {
		return false
	}
	if b.board[x][y].GetState() != STATE_VOID {
		return false
	}
	b.SetGO(x, y, state)
	return true
}

func (b *Board) ShowBoard() string {
	var buffer bytes.Buffer
	for j := 0; j < _length; j++ {
		for i := 0; i < _width; i++ {
			switch b.board[i][j].GetState() {
			case STATE_BLACK:
				buffer.WriteByte('@')
				break
			case STATE_WHITE:
				buffer.WriteByte('O')
				break
			case STATE_VOID:
				buffer.WriteByte('+')
				break
			default:
				break
			}
			buffer.WriteString(" ")
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

// func main() {
// 	board := NewBoard()
// 	exit := false
// 	state := STATE_BLACK
// 	x := 0
// 	y := 0

// 	for !exit {
// 		fmt.Println(board.ShowBoard())
// 		switch state {
// 		case STATE_BLACK:
// 			fmt.Println("it's BLACK's turn")
// 			break
// 		case STATE_WHITE:
// 			fmt.Println("it's WHITE's turn")
// 			break
// 		}
// 		fmt.Println("please int x:")
// 		fmt.Scanln(&x)
// 		x = x - 1
// 		fmt.Println("please int y:")
// 		fmt.Scanln(&y)
// 		y = y - 1
// 		if x < 0 || x >= _width || y < 0 || y >= _length {
// 			fmt.Println("wrong input")
// 			continue
// 		}

// 		if board.Drop(x, y, state) {
// 			// board.board[x][y].state = state
// 			board.Eat(x, y)
// 			switch state {
// 			case STATE_BLACK:
// 				state = STATE_WHITE
// 				break
// 			case STATE_WHITE:
// 				state = STATE_BLACK
// 				break
// 			}
// 		} else {
// 			fmt.Println("can not drop here")
// 		}
// 		// fmt.Println(board.board[x][y].state)
// 	}

// }
