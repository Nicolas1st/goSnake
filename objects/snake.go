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

	// adding new head
	snake.blocks = append(snake.blocks, newHead)

	switch board.grid[x][y] {
	case snake.bodySymbol:
		tail := snake.blocks[0]
		// removing the tail part
		snake.blocks = snake.blocks[1:]
		// if it's the tail, then it won't be there on the next iteration
		// so it's not a collision
		if !(tail.x == x && tail.y == y) {
			return false
		}
	case board.emptyCellSymbol:
		// simple movement
		// removing the tail part
		// the head was already added
		snake.blocks = snake.blocks[1:]
	case food.Symbol:
		// the tail part is not removed because
		// eating food increases snake's length
		// (head was already added, so it's +1 already)
		food.setFreeCoordinates(board)
	}

	return true

}

func (snake *Snake) SetDirection(direction string) {
	if oppositeDirections[snake.direction] != direction {
		snake.direction = direction
	}
}

func (snake *Snake) GetSnakeLength() int {
	return len(snake.blocks)
}

func (snake *Snake) renderOnBoard(board *Board) {

	// disaplaying the rest of the body
	for _, block := range snake.blocks[:len(snake.blocks)-1] {
		board.grid[block.x][block.y] = snake.bodySymbol
	}

	// displaying the head
	var head snakeBlock = snake.blocks[len(snake.blocks)-1]
	board.grid[head.x][head.y] = snake.headSymbol

}
