# Snake Game In Go

## My First Ever Go Project

<img src="./goSnakeGameplay.gif" width="500" height="500"/>

## Contents
  - [Structure](#structure)
  - [Installation And Running](#installation-and-running)
  - [Controls](#controls)
  - [Configuration](#configuration)

## Structure

- [main.go](./main.go)

  Entry point of the game

- [objects](./objects)

  All of the game objects + corresponding test files

  - [board.go](./objects/board.go)

    contains the object that renders all other objects (snake.go and food.go)

    that have the method `renderOnBoard`

  - [keymap.go](./objects/keymap.go)
    
    contains hashmaps

    - keyMap

      pressed keys to snake movements on the board

    - oppositeKeys

      opposite snake directions

- [config](./config)

  Config objects and a toml config parsing function

- [color](./color)

  Helper package for adding color to the terminal + tests

## Installation And Running

- Make sure you have the version of Go specified in [go.mod](./go.mod)

- Clone the repository:

```sh
git clone github.com/Nicolas1st/goSnake
```

- Download the dependencies:

```sh
go mod download
```

- Go to the directory (folder) that contains the project's

  files and run the following command in the terminal:

```sh
go run main.go
```

## Controls

- To control the snake use 'w', 'a', 's', 'd' keys

  to go up, left, down, and right

  The movement keys can be changed [here](#configuration)

  To quit the game press 'q'
    
## Configuration

The [config.toml](./config.toml) can be used.

Set the environment variable to the config file name,

other configs can be used as well.

On Linux:

```sh
export gosnake_config_file=config.toml
```

On Windows:

```cmd
set gosnake_config_file=config.toml
```

Then run:

```sh
go run main.go
```

To change controls, edit the [keymap.go](./objects/keymap.go):
```go
var keyMap = map[rune]string{
	'w': "up",
	'a': "left",
	's': "down",
	'd': "right",
}
```
