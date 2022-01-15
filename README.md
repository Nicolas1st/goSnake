# Snake Game In GO

Simple snake game (My first ever GO project)

- [Snake Game In GO](#snake-game-in-go)
  - [Structure](#structure)
  - [Configuration](#configuration)
  - [Installation And Running](#installation-and-running)

## Structure

- main.go

  Entry point of the game

- objects

  All of the game objects + corresponding test files

  - board.go

    contains the object that renders all other objects (snake.go and food.go)

    that have the method `renderOnBoard`

  - keymap
    
    contains hashmaps

    - keyMap

      pressed keys to snake movements on the board

    - oppsiteKeys

      opposite snake directions

- color

  Helper package for adding color to the terminal + tests
    
## Configuration

```go
var board objects.Board = objects.CreateBoard(20, 20, " ")
var snake = objects.CreateSnake(color.TextToGreen("$"), color.TextToGreen("#"), 10, 10, "up")
var food = objects.CreateFood(15, 15, color.TextToRed("@"))
```
The single character values define what characters will be used

to display corresponding objects, values define

board size or snake initial position,

see more [here](./main.go)

```go
var keyMap = map[rune]string{
	'w': "up",
	'a': "left",
	's': "down",
	'd': "right",
}
```
Can be configured [here](./objects/keymap.go)

## Installation And Running

- Make sure you have the version of GO specified in [go.mod](./go.mod)

- Go to the directory (folder) that contains projects files and run the following command in the terminal

```shell
go run main.go
```

- To control the snake use 'w', 'a', 's', 'd' keys,

  keys can be recongifured as shown [here](#configuration)

- To quit the game press 'q'