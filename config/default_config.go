package config

var defaultConfig Config = Config{
	ScoreToWin:       30,
	TimePerFrameInMs: 100,

	Board: &BoardConfig{
		Height:          20,
		Width:           20,
		EmptyTileSymbol: " ",
	},

	Snake: &SnakeConfig{
		HeadColor:  "yellow",
		HeadSymbol: "$",

		BodyColor:  "green",
		BodySymbol: "#",

		X: 10,
		Y: 10,

		MovementDirection: "right",
	},

	Food: &FoodConfig{
		Color:  "red",
		Symbol: "@",

		X: 15,
		Y: 15,
	},

	ActionToKeyMap: &Keyboard{
		Up:    "w",
		Left:  "a",
		Down:  "s",
		Right: "d",

		Quit: "q",
	},

	KeyToActionMap: map[string]string{
		"w": "up",
		"a": "left",
		"s": "down",
		"d": "right",

		"q": "quit",
	},
}
