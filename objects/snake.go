package objects

type snakeBlock struct {
	x, y int
}

type Snake struct {
	headSymbol string
	bodySymbol string
	blocks     []snakeBlock
	direction  string
}

func CreateSnake(headSymbol string, bodySymbol string, x, y int, movement string) Snake {
	var head snakeBlock = snakeBlock{x, y}
	return Snake{headSymbol, bodySymbol, []snakeBlock{head}, movement}
}

// Move returns true if the move made was valid, false otherwise
func (snake *Snake) Move(board *Board, food *Food) bool {

	var oldHead = snake.blocks[len(snake.blocks)-1]
	var x, y int = oldHead.x, oldHead.y

	switch snake.direction {
	case "up":
		x = (x - 1 + board.width) % board.width
		y = y % board.height
	case "down":
		x = (x + 1 + board.width) % board.width
		y = y % board.height
	case "left":
		x = x % board.width
		y = (y - 1 + board.height) % board.height
	case "right":
		x = x % board.width
		y = (y + 1 + board.height) % board.height
	}

	var newHead snakeBlock = snakeBlock{x: x, y: y}

	snake.blocks = append(snake.blocks, newHead)

	switch board.grid[x][y] {
	case snake.bodySymbol:
		return false
	case board.emptyCellSymbol:
		// removing the tail part
		snake.blocks = snake.blocks[1:]
	case food.Symbol:
		food.setFreeCoordinates(board, snake)
	}

	return true

}

func (snake *Snake) SetDirection(pressedKey rune) {
	direction, ok := keyMap[pressedKey]
	if ok && oppositeKeys[snake.direction] != direction {
		snake.direction = direction
	}
}

func (snake *Snake) GetSnakeLength() int {
	return len(snake.blocks)
}

func (snake *Snake) renderOnBoard(board *Board) {

	// displaying the head
	var head snakeBlock = snake.blocks[len(snake.blocks)-1]

	board.grid[head.x][head.y] = snake.headSymbol

	// disaplaying the rest of the body
	for _, block := range snake.blocks[:len(snake.blocks)-1] {
		board.grid[block.x][block.y] = snake.bodySymbol
	}

}
