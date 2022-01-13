package color

var red = "\033[31m"
var green = "\033[32m"

func TextToGreen(text string) string {
	return green + text + "\033[0m"
}

func TextToRed(text string) string {
	return red + text + "\033[0m"
}
