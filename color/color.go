package color

type Color string

const (
	Red                 Color = "\033[31m"
	Green               Color = "\033[32m"
	Yellow              Color = "\033[33m"
	endColoringSequence Color = "\033[0m"
)

func TextToColor(text string, color Color) string {
	return string(color) + text + string(endColoringSequence)
}
