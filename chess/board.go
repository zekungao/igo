package main

import (
	"bytes"
	"fmt"
)

const _length int = 19
const _width int = 19

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

type Grid struct {
	x     int
	y     int
	state State
}

func NewGrid(mx int, my int) *Grid {
	return &Grid{x: mx, y: my, state: STATE_VOID}
}

func (this *Grid) SetState(s State) {
	this.state = s
}

func (this *Grid) ClearState() {
	this.state = STATE_VOID
}

func (this *Grid) GetState() State {
	return this.state
}

func (this *Grid) EqualState(s State) bool {
	if this.state == s {
		return true
	}
	return false
}

type Board struct {
	originX int
	originY int
	board   [_length][_width]Grid
}

func NewBoard() *Board {
	chessBoard := [_length][_width]Grid{}
	for i := 0; i < _width; i++ {
		for j := 0; j < _length; j++ {
			chessBoard[i][j] = *NewGrid(i, j)
		}
	}
	return &Board{originX: -1, originY: -1, board: chessBoard}
}

func (this *Board) SetGO(x int, y int, state State) {
	// fmt.Println(this.board[x][y].state)
	this.board[x][y].SetState(state)
	// fmt.Println(this.board[x][y].state)
}

func (this *Board) JudgeDA(x int, y int, originState State, ban BanVector) bool {
	if x == this.originX && y == this.originY {
		return false
	}
	if this.board[x][y].EqualState(STATE_VOID) {
		return true
	}
	if !this.board[x][y].EqualState(originState) {
		return false
	}
	return this.DeathAlive(x, y, originState, ban)
}

func (this *Board) DeathAlive(x int, y int, originState State, ban BanVector) bool {
	//true:alive, false:death
	if x != 0 && ban != BAN_LEFT {
		if this.JudgeDA(x-1, y, originState, BAN_RIGHT) {
			return true
		}
	}
	if x != _width-1 && ban != BAN_RIGHT {
		if this.JudgeDA(x+1, y, originState, BAN_LEFT) {
			return true
		}
	}
	if y != 0 && ban != BAN_UP {
		if this.JudgeDA(x, y-1, originState, BAN_DOWN) {
			return true
		}
	}
	if y != _length-1 && ban != BAN_DOWN {
		if this.JudgeDA(x, y+1, originState, BAN_UP) {
			return true
		}
	}
	return false
}

func (this *Board) ResetOrigin() {
	this.originX = -1
	this.originY = -1
}

func (this *Board) CanEat(x int, y int, originState State) bool {
	this.originX = x
	this.originY = y
	if x != 0 && !this.board[x-1][y].EqualState(originState) {
		if !this.DeathAlive(x-1, y, this.board[x-1][y].GetState(), BAN_RIGHT) {
			this.ResetOrigin()
			return true
		}
	}
	if x != _width-1 && !this.board[x+1][y].EqualState(originState) {
		if !this.DeathAlive(x+1, y, this.board[x+1][y].GetState(), BAN_LEFT) {
			this.ResetOrigin()
			return true
		}
	}
	if y != 0 && !this.board[x][y-1].EqualState(originState) {
		if !this.DeathAlive(x, y-1, this.board[x][y-1].GetState(), BAN_DOWN) {
			this.ResetOrigin()
			return true
		}
	}
	if y != _length-1 && !this.board[x][y+1].EqualState(originState) {
		if !this.DeathAlive(x, y+1, this.board[x][y+1].GetState(), BAN_UP) {
			this.ResetOrigin()
			return true
		}
	}
	this.ResetOrigin()
	return false
}

func (this *Board) CanDrop(x int, y int, originState State) bool {
	if this.CanEat(x, y, originState) {
		return true
	}
	if this.DeathAlive(x, y, originState, BAN_VOID) {
		return true
	}
	return false
}

func (this *Board) Eating(x int, y int, ban BanVector) {
	originState := this.board[x][y].GetState()

	if x != 0 && this.board[x-1][y].EqualState(originState) && ban != BAN_LEFT {
		this.Eating(x-1, y, BAN_RIGHT)
	}
	if x != _width-1 && this.board[x+1][y].EqualState(originState) && ban != BAN_RIGHT {
		this.Eating(x+1, y, BAN_LEFT)
	}
	if y != 0 && this.board[x][y-1].EqualState(originState) && ban != BAN_DOWN {
		this.Eating(x, y-1, BAN_UP)
	}
	if y != _length-1 && this.board[x][y+1].EqualState(originState) && ban != BAN_UP {
		this.Eating(x, y+1, BAN_DOWN)
	}
	this.board[x][y].ClearState()
}

func (this *Board) Eat(x int, y int) {
	originState := this.board[x][y].GetState()
	if !this.CanEat(x, y, originState) {
		fmt.Println("not eat")
		return
	}
	if x != 0 && !this.board[x-1][y].EqualState(originState) {
		if !this.DeathAlive(x-1, y, this.board[x-1][y].GetState(), BAN_RIGHT) {
			this.Eating(x-1, y, BAN_RIGHT)
		}
	}
	if x != _width-1 && !this.board[x+1][y].EqualState(originState) {
		if !this.DeathAlive(x+1, y, this.board[x+1][y].GetState(), BAN_LEFT) {
			this.Eating(x+1, y, BAN_LEFT)
		}
	}
	if y != 0 && !this.board[x][y-1].EqualState(originState) {
		if !this.DeathAlive(x, y-1, this.board[x][y-1].GetState(), BAN_DOWN) {
			this.Eating(x, y-1, BAN_UP)
		}
	}
	if y != _length-1 && !this.board[x][y+1].EqualState(originState) {
		if !this.DeathAlive(x, y+1, this.board[x][y+1].GetState(), BAN_UP) {
			this.Eating(x, y+1, BAN_DOWN)
		}
	}
}

func (this *Board) Drop(x int, y int, state State) bool {
	if !this.CanDrop(x, y, state) {
		return false
	}
	if this.board[x][y].GetState() != STATE_VOID {
		return false
	}
	this.SetGO(x, y, state)
	return true
}

func (this *Board) ShowBoard() string {
	var buffer bytes.Buffer
	for j := 0; j < _length; j++ {
		for i := 0; i < _width; i++ {
			switch this.board[i][j].GetState() {
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

func main() {
	board := NewBoard()
	exit := false
	state := STATE_BLACK
	x := 0
	y := 0

	for !exit {
		fmt.Println(board.ShowBoard())
		switch state {
		case STATE_BLACK:
			fmt.Println("it's BLACK's turn")
			break
		case STATE_WHITE:
			fmt.Println("it's WHITE's turn")
			break
		}
		fmt.Println("please int x:")
		fmt.Scanln(&x)
		x = x - 1
		fmt.Println("please int y:")
		fmt.Scanln(&y)
		y = y - 1
		if x < 0 || x >= _width || y < 0 || y >= _length {
			fmt.Println("wrong input")
			continue
		}

		if board.Drop(x, y, state) {
			// board.board[x][y].state = state
			board.Eat(x, y)
			switch state {
			case STATE_BLACK:
				state = STATE_WHITE
				break
			case STATE_WHITE:
				state = STATE_BLACK
				break
			}
		} else {
			fmt.Println("can not drop here")
		}
		// fmt.Println(board.board[x][y].state)
	}

}
