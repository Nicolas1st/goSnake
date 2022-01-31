package objects

var keyMap = map[rune]string{
	'w': "up",
	'a': "left",
	's': "down",
	'd': "right",
}

var oppositeKeys = map[string]string{
	"up":    "down",
	"down":  "up",
	"left":  "right",
	"right": "left",
}
