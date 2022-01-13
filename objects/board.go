package objects

import (
	"fmt"
	"os"
	"os/exec"
)

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

func (board *Board) Rerender(objects ...renderableOnBoard) {

	// replacing the old board with a clean one
	board.grid = createGrid(board.height, board.width, board.emptyCellSymbol)

	for _, object := range objects {
		object.renderOnBoard(board)
	}

	// clearing the terminal from the last displayed grid
	command := exec.Command("clear")
	command.Stdout = os.Stdout
	err := command.Run()
	if err != nil {
		command = exec.Command("cls")
		command.Stdout = os.Stdout
		command.Run()
	}

	for _, row := range board.grid {
		fmt.Println(row)
	}

}
