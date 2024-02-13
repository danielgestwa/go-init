package main

import "testing"

func TestPerimiter(t *testing.T) {
	rect := Rectangle{10.0, 20.0}
	expected := rect.Perimiter()
	want := 60.0

	if expected != want {
		t.Errorf("Expected %.2f, got %.2f", expected, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%q: got %g want %g", tt.name, got, tt.want)
			}
		})
	}

}
