package config

var DefaultConfig Config = Config{
	ScoreToWin: 30,

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
}