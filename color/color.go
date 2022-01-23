package color

var (
	red                 = "\033[31m"
	green               = "\033[32m"
	yellow              = "\033[33m"
	endColoringSequence = "\033[0m"
)

func TextToGreen(text string) string {
	return green + text + endColoringSequence
}

func TextToRed(text string) string {
	return red + text + endColoringSequence
}

func TextToYellow(text string) string {
	return yellow + text + endColoringSequence
}
