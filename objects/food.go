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
func (food *Food) setFreeCoordinates(board *Board, snake *Snake) {

	var x, y int
	var snakeHead snakeBlock = snake.blocks[len(snake.blocks)-1]

	// if the field fills up it might result in the infinte loop
	// figure out how to deal with it
	for {

		// should not appear at the position of head
		// not checking the tail position because the food
		// is being generated only when the old one has been eaten
		// since that block is occupied
		if x == snakeHead.x && y == snakeHead.y {
			continue
		}

		x, y = rand.Intn(board.width), rand.Intn(board.height)

		if board.grid[x][y] == board.emptyCellSymbol {
			break
		}
	}

	food.x = x
	food.y = y

}

func (food *Food) renderOnBoard(board *Board) {
	board.grid[food.x][food.y] = food.Symbol
}
