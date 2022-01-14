package objects

import "testing"

func TestCreateFood(t *testing.T) {

	var createFoodTests = []struct {
		x, y       int
		foodSymbol string
		expected   Food
	}{
		{1, 2, "@", Food{1, 2, "@"}},
		{3, 1, "x", Food{3, 1, "x"}},
		{4, 2, "o", Food{4, 2, "o"}},
	}

	for _, test := range createFoodTests {
		if output := CreateFood(test.x, test.y, test.foodSymbol); output != test.expected {
			t.Errorf("Expected %q, received %q", test.expected, output)
		}
	}

}
