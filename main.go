package main

import (
	"fmt"
	"log"
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
		success := snake.Move(&board, &food)
		if !success {
			fmt.Printf("Game Over %v\n", snake.GetSnakeLength())
			return
		}
		board.Rerender(&snake, &food)
		fmt.Printf("Score: %v\n", snake.GetSnakeLength())
	}

}
