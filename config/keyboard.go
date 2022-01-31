package config

type Keyboard struct {
	Up    string
	Left  string
	Down  string
	Right string

	Quit string
}

func (keys *Keyboard) toSwappedKeysAndValuesMap() map[string]string {
	return map[string]string{
		keys.Up:    "up",
		keys.Down:  "down",
		keys.Left:  "left",
		keys.Right: "right",

		keys.Quit: "quit",
	}
}
