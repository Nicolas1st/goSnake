package color

import "testing"

func testColoring(t *testing.T, color Color, f func(string, Color) string) {

	// a copy of the function being tested,
	// if the exported function changes,
	// it will be known
	createExpectedOutput := func(text string, color Color) string {
		endColoringSequence := "\033[0m"
		return string(color) + text + string(endColoringSequence)
	}

	tests := []struct{ input, expected string }{
		{"a", createExpectedOutput("a", color)},
		{"b", createExpectedOutput("b", color)},
		{"c", createExpectedOutput("c", color)},
	}

	for _, test := range tests {
		output := f(test.input, color)
		if output != test.expected {
			t.Errorf("Expected %q, received %q", test.expected, output)
		}
	}

}

func TestTextToColor(t *testing.T) {

	colors := []Color{
		Green,
		Red,
		Yellow,
	}

	for _, color := range colors {
		testColoring(t, color, TextToColor)
	}

}
