package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/Nicolas1st/goSnake/color"
	"github.com/Nicolas1st/goSnake/config"
	"github.com/Nicolas1st/goSnake/objects"
	"github.com/mattn/go-tty"
)

func main() {

	// reading the config
	configName, _ := os.LookupEnv("gosnake_config_file")
	config := config.NewConfig(configName)

	// initializing objects
	var board objects.Board = objects.CreateBoard(
		config.Board.Height,
		config.Board.Width,
		config.Board.EmptyTileSymbol,
	)

	var snake = objects.CreateSnake(
		color.TextToColor(
			config.Snake.HeadSymbol,
			config.Snake.HeadColor,
		),
		color.TextToColor(
			config.Snake.BodySymbol,
			config.Snake.BodyColor,
		),
		config.Snake.X,
		config.Snake.Y,
		config.Snake.MovementDirection,
	)

	var food = objects.CreateFood(
		config.Food.X,
		config.Food.Y,
		color.TextToColor(
			config.Food.Symbol,
			config.Food.Color,
		),
	)

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

		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt)

		<-signalChannel
		inputChannel <- 'q'

	}(inputChannel)

	// game loop
	for {

		// time between game frames
		time.Sleep(time.Duration(config.TimePerFrameInMs) * time.Millisecond)

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
		if snakeLength == config.ScoreToWin {
			fmt.Printf(color.TextToColor("You win! Your snake reached the length of %v\n", "green"), config.ScoreToWin)
			return
		} else if snakeDidNotHitItself {
			fmt.Printf(color.TextToColor("Score: %v\n", "yellow"), snakeLength)
		} else {
			fmt.Printf(color.TextToColor("Game Over. Score: %v\n", "red"), snakeLength)
			return
		}

	}

}
