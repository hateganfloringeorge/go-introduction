package structs

import (
	"testing"
)

func TestStructs(t *testing.T) {
	
	areaTest := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 7.5,Height: 4.0}, hasArea: 30.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12,Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

	t.Run("perimeter of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.5, 2.0}
		
		result := Perimeter(rectangle)
		expected := 25.0

		if expected != result {
			t.Errorf("expected %g but got %g", expected, result)
		}
	})
}