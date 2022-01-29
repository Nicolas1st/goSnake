package color

type Color string

const (
	Red                 Color = "\033[31m"
	Green               Color = "\033[32m"
	Yellow              Color = "\033[33m"
	endColoringSequence Color = "\033[0m"
)

var nameToColor = map[string]Color{
	"red":    Red,
	"green":  Green,
	"yellow": Yellow,
}

func TextToColor(text string, colorName string) string {
	return string(nameToColor[colorName]) + text + string(endColoringSequence)
}
