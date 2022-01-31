package routines

import (
	"log"

	"github.com/mattn/go-tty"
)

func CaptureUserInput(tty *tty.TTY, userInputChannel chan rune) {
	for {
		pressedKey, err := tty.ReadRune()

		if err != nil {
			log.Fatal(err)
		}

		userInputChannel <- pressedKey
	}
}
