package color

import "testing"

func testColoring(t *testing.T, colorName string, f func(string, string) string) {

	// a copy of the function being tested,
	// if the exported function changes,
	// it will be known
	createExpectedOutput := func(text, colorName string) string {
		endColoringSequence := "\033[0m"
		return string(nameToColor[colorName]) + text + string(endColoringSequence)
	}

	tests := []struct{ input, expected string }{
		{"a", createExpectedOutput("a", colorName)},
		{"b", createExpectedOutput("b", colorName)},
		{"c", createExpectedOutput("c", colorName)},
	}

	for _, test := range tests {
		output := f(test.input, colorName)
		if output != test.expected {
			t.Errorf("Expected %q, received %q", test.expected, output)
		}
	}

}

func TestTextToColor(t *testing.T) {

	colors := []string{
		"green",
		"red",
		"yellow",
	}

	for _, color := range colors {
		testColoring(t, color, TextToColor)
	}

}
