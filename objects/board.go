package objects

import "fmt"

type Board struct {
	height, width   int
	grid            [][]string
	emptyCellSymbol string
}

func createGrid(height, width int, symbol string) [][]string {

	grid := make([][]string, height)

	row := make([]string, width)
	for i := range row {
		row[i] = symbol
	}

	for i := range grid {
		tmp := make([]string, len(row))
		copy(tmp, row)
		grid[i] = tmp
	}

	return grid

}

// CreateBoard creates a 2d array used to display the field of the snake game
func CreateBoard(height, width int, emtpyCellSymbol string) Board {
	return Board{
		height:          height,
		width:           width,
		grid:            createGrid(height, width, emtpyCellSymbol),
		emptyCellSymbol: emtpyCellSymbol,
	}
}

func (board *Board) Render(objects ...renderableOnBoard) {

	// replacing the old board with a clean one
	board.grid = createGrid(board.height, board.width, board.emptyCellSymbol)

	for _, object := range objects {
		object.renderOnBoard(board)
	}

	for _, row := range board.grid {
		fmt.Println(row)
	}

}
