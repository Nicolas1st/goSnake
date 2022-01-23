package color

import "testing"

func testColoring(t *testing.T, color string, f func(string) string) {

	createExpectedOutput := func(text, color string) string {

		endColoringSequence := "\033[0m"
		return color + text + endColoringSequence

	}

	tests := []struct{ input, expected string }{
		{"a", createExpectedOutput("a", color)},
		{"b", createExpectedOutput("b", color)},
		{"c", createExpectedOutput("c", color)},
	}

	for _, test := range tests {
		output := f(test.input)
		if output != test.expected {
			t.Errorf("Expected %q, received %q", test.expected, output)
		}
	}

}

func TestTextToGreen(t *testing.T) {

	greenColor := "\033[32m"
	testColoring(t, greenColor, TextToGreen)

}

func TestTextToRed(t *testing.T) {

	redColor := "\033[31m"
	testColoring(t, redColor, TextToRed)

}

func TestTextToYellow(t *testing.T) {

	yellowColor := "\033[33m"
	testColoring(t, yellowColor, TextToYellow)

}
