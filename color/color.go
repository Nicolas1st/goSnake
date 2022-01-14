package color

var (
	red                 = "\033[31m"
	green               = "\033[32m"
	endColoringSequence = "\033[0m"
)

func TextToGreen(text string) string {
	return green + text + endColoringSequence
}

func TextToRed(text string) string {
	return red + text + endColoringSequence
}
