package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestAreaTable(t *testing.T) {

	// "anonymous struct"
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		// {Rectangle{12, 6}, 72.0},
		// {Circle{10}, 314.1592653589793},
		// {Triangle{12, 6}, 36.0},

		// field naming: optional
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}

}
