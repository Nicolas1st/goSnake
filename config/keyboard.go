package config

type KeyboardKeys struct {
	Up    rune
	Left  rune
	Down  rune
	Right rune

	Quit rune
}

func (keys *KeyboardKeys) toMap() map[rune]string {
	return map[rune]string{
		keys.Up:    "up",
		keys.Down:  "down",
		keys.Left:  "left",
		keys.Right: "right",

		keys.Quit: "quit",
	}
}
