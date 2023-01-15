package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{3, 4}
	result := Perimeter(rectangle)
	expected := 14.0

	if expected != result {
		t.Errorf("Expected '%.2f' but received '%.2f' when Width = '%.2f', Height = '%.2f'", expected, result, rectangle.Width, rectangle.Height)
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		//{shape: Rectangle{4, 5}, expected: 20.0},
		//{shape: Circle{10}, expected: 314.1592653589793},
		{"Rectangle", Rectangle{4, 5}, 20.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36},
	}

	for _, testCase := range testCases {
		result := testCase.shape.Area()

		if testCase.expected != result {
			t.Errorf("Expected '%.2f' but received '%.2f' when shape = '%#v'", testCase.expected, result, testCase.shape)
		}
	}
}
