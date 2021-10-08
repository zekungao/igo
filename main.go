package main

import (
	"fmt"
	"igo/src"
)

func main() {
	board := src.NewBoard()
	exit := false
	state := src.STATE_BLACK
	x := 0
	y := 0

	for !exit {
		fmt.Println(board.ShowBoard())
		switch state {
		case src.STATE_BLACK:
			fmt.Println("it's BLACK's turn")
		case src.STATE_WHITE:
			fmt.Println("it's WHITE's turn")
		}
		fmt.Println("please int x:")
		fmt.Scanln(&x)
		x = x - 1
		fmt.Println("please int y:")
		fmt.Scanln(&y)
		y = y - 1
		if x < 0 || x >= board.Width || y < 0 || y >= board.Length {
			fmt.Println("wrong input")
			continue
		}

		if board.Drop(x, y, state) {
			board.Eat(x, y)
			switch state {
			case src.STATE_BLACK:
				state = src.STATE_WHITE
			case src.STATE_WHITE:
				state = src.STATE_BLACK
			}
		} else {
			fmt.Println("can not drop here")
		}
	}

}
