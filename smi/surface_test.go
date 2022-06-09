package smi

import "testing"

var assertPerimeter = func(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got: %.2f want %.2f", got, want)
	}
}

var assertArea = func(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	// t.Run("calculate perimeter", func(t *testing.T){
	// 	got:=Perimeter(10.0,10.0)
	// 	want:=40.0

	// 	assertPerimeter(t, got, want)
	// })
	t.Run("calculate perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0
		assertPerimeter(t, got, want)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}
	areaCases := []struct {
		name string
		shape Shape
		hasArea  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 64.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}
	// t.Run("calculate area", func(t *testing.T) {
	// 	got := Area(12.0, 6.0)
	// 	want := 72.0

	// 	assertArea(t, got, want)
	// })
	// t.Run("calculate area of rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t, rectangle, 72.0)
	// })
	// t.Run("calculate area of circle", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })
	
	for _, test := range areaCases {
		t.Run(test.name, func(t *testing.T){
			checkArea(t, test.shape, test.hasArea)
		})
	}

}
