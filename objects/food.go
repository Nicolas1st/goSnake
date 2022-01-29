package objects

import (
	"math/rand"
)

type Food struct {
	x, y   int
	Symbol string
}

func CreateFood(x, y int, foodSymbol string) Food {
	return Food{x: x, y: y, Symbol: foodSymbol}
}

// uses the board.grid from the previous iteration
func (food *Food) setFreeCoordinates(board *Board) {

	var newX, newY int

	for {

		// should not appear at the same place,
		// because snake's head is located there
		if newX == food.x && newY == food.y {
			continue
		}

		newX, newY = rand.Intn(board.width), rand.Intn(board.height)

		if board.grid[newX][newY] == board.emptyCellSymbol {
			food.x = newX
			food.y = newY
			break
		}

	}

}

func (food *Food) renderOnBoard(board *Board) {
	board.grid[food.x][food.y] = food.Symbol
}
