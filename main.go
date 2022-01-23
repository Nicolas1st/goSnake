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

func main() {

	var board objects.Board = objects.CreateBoard(20, 20, " ")
	var snake = objects.CreateSnake(color.TextToGreen("$"), color.TextToGreen("#"), 10, 10, "up")
	var food = objects.CreateFood(15, 15, color.TextToRed("@"))

	inputChannel := make(chan rune)

	var pressedKey rune
	go func(channel chan rune) {
		tty, err := tty.Open()

		if err != nil {
			log.Fatal(err)
		}

		for {
			pressedKey, err := tty.ReadRune()
			if pressedKey == 'q' {
				tty.Close()
			}

			if err != nil {
				log.Fatal(err)
			}

			channel <- pressedKey
		}
	}(inputChannel)

	for {
		time.Sleep(100 * time.Millisecond)
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

		board.Render(&snake, &food)

		snakeDidNotHitItself := snake.Move(&board, &food)
		snakeLength := snake.GetSnakeLength()

		if snakeLength == scoreToWin {
			fmt.Printf(color.TextToGreen("You win! Your snake reached the lenght of %v\n"), snakeLength)
			return
		} else if snakeDidNotHitItself {
			fmt.Printf(color.TextToYellow("Score: %v\n"), snakeLength)
		} else {
			fmt.Printf(color.TextToRed("Game Over %v\n"), snakeLength)
			return
		}

	}

}
