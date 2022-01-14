package objects

import "testing"

func TestCreateSnake(t *testing.T) {

	var createFoodTests = []struct {
		headSymbol, bodySymbol string
		x, y                   int
		direction              string
		expected               Snake
	}{
		{"@", "x", 1, 2, "up", Snake{"@", "x", []snakeBlock{{1, 2}}, "up"}},
		{"o", ".", 2, 3, "down", Snake{"o", ".", []snakeBlock{{2, 3}}, "down"}},
		{"p", "a", 3, 4, "left", Snake{"p", "a", []snakeBlock{{3, 4}}, "left"}},
		{"w", "l", 4, 5, "right", Snake{"w", "l", []snakeBlock{{4, 5}}, "right"}},
	}

	for _, test := range createFoodTests {

		output := CreateSnake(test.headSymbol, test.bodySymbol, test.x, test.y, test.direction)

		if output.headSymbol != test.expected.headSymbol {
			t.Errorf("Expected  headSymbol equal to %q, received %q", test.expected.headSymbol, output.headSymbol)
		}

		if output.bodySymbol != test.expected.bodySymbol {
			t.Errorf("Expected  bodySymbol equal to %q, received %q", test.expected.bodySymbol, output.bodySymbol)
		}

		if output.direction != test.expected.direction {
			t.Errorf("Expected direction equal to %q, received %q", test.direction, output.direction)
		}

		if l1, l2 := len(output.blocks), len(test.expected.blocks); l1 != l2 {
			t.Errorf("Expected the snake to have the lenght of %q, received %q", l1, l2)
			t.FailNow()
		}

		for i := range output.blocks {
			receivedBlock, expectedBlock := output.blocks[i], test.expected.blocks[i]

			if receivedBlock.x != expectedBlock.x {
				t.Errorf("Expected the block %q to have equal x coordinates, received %q, expected %q",
					i, receivedBlock.x, expectedBlock.x)
			}

			if receivedBlock.y != expectedBlock.y {
				t.Errorf("Expected the block %q to have equal y coordinates, received %q, expected %q",
					i, receivedBlock.y, expectedBlock.y)
			}
		}
	}

}
