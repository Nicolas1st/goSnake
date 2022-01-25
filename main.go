package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/Nicolas1st/goSnake/color"
	"github.com/Nicolas1st/goSnake/objects"
	"github.com/mattn/go-tty"
)

func main() {

	const scoreToWin int = 100

	// initializing objects
	var board objects.Board = objects.CreateBoard(20, 20, " ")
	var snake = objects.CreateSnake(color.TextToYellow("$"), color.TextToGreen("#"), 10, 10, "up")
	var food = objects.CreateFood(15, 15, color.TextToRed("@"))

	// opening a tty to read input from
	t, err := tty.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer t.Close()

	// running a go routine to get user input
	inputChannel := make(chan rune)
	var pressedKey rune
	go func(tty *tty.TTY, inputChannel chan rune) {

		for {

			pressedKey, err := tty.ReadRune()

			if err != nil {
				log.Fatal(err)
			}

			inputChannel <- pressedKey

		}

	}(t, inputChannel)

	// making sure the tty is realeased even if forceful quit occurs
	go func(inputChannel chan rune) {

		signalChannel := make(chan os.Signal)
		signal.Notify(signalChannel, os.Interrupt)

		<-signalChannel
		inputChannel <- 'q'

	}(inputChannel)

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

		// clearing the old frame
		command := exec.Command("clear")
		command.Stdout = os.Stdout

		err := command.Run()
		if err != nil {
			command = exec.Command("cls")
			command.Stdout = os.Stdout
			command.Run()
		}

		// updating the objects before displaying the new board
		snakeDidNotHitItself := snake.Move(&board, &food)
		snakeLength := snake.GetSnakeLength()

		board.Render(&snake, &food)

		// in-game messages
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
