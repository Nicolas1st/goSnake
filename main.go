package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/Nicolas1st/goSnake/color"
	"github.com/Nicolas1st/goSnake/objects"
	"github.com/mattn/go-tty"
)

func clearScreen() {

	command := exec.Command("clear")
	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		command = exec.Command("cls")
		command.Stdout = os.Stdout
		command.Run()
	}

}

func readUserInput(tty *tty.TTY, inputChannel chan rune) {

	for {

		pressedKey, err := tty.ReadRune()

		if err != nil {
			log.Fatal(err)
		}

		inputChannel <- pressedKey

	}

}

func main() {

	const scoreToWin int = 100

	// initializing objects
	var board objects.Board = objects.CreateBoard(20, 20, " ")
	var snake = objects.CreateSnake(color.TextToYellow("$"), color.TextToGreen("#"), 10, 10, "up")
	var food = objects.CreateFood(15, 15, color.TextToRed("@"))

	// opening a tty to read input from
	tty, err := tty.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer tty.Close()

	// running a go routine to get user input
	inputChannel := make(chan rune)
	var pressedKey rune
	go readUserInput(tty, inputChannel)

	// game loop
	for {

		// time between game frames
		time.Sleep(100 * time.Millisecond)

		// getting user input if it there was any
		select {
		case pressedKey = <-inputChannel:
			if pressedKey == 'q' {
				return
			}
			snake.SetDirection(pressedKey)
		default:
		}

		// rendering a new frame, the old one should be cleared
		clearScreen()

		snakeDidNotHitItself := snake.Move(&board, &food)
		snakeLength := snake.GetSnakeLength()

		board.Render(&snake, &food)

		if snakeLength == scoreToWin {
			fmt.Printf(color.TextToGreen("You win! Your snake reached the length of %v\n"), scoreToWin)
			return
		} else if snakeDidNotHitItself {
			fmt.Printf(color.TextToYellow("Score: %v\n"), snakeLength)
		} else {
			fmt.Printf(color.TextToRed("Game Over. Score: %v\n"), snakeLength)
			return
		}

	}

}
