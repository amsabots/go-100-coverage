package structs

import "testing"


func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{width: 10.0, height: 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}


func TestArea(t *testing.T) {

	testCases := []struct{
		s Shape
		want float64
		name string
	}{
		{
            s: Rectangle{12, 6},
			want: 72.0,
			name: "Rectangle",
		},
		{
           s: Circle{10},
		   want: 314.1592653589793,
		   name: "Circle",
		},
	}

	for _, tt := range testCases{
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Area()
		if got != tt.want {
			t.Errorf("%#v got %.2f want %.2f", tt.s, got, tt.want)
		}
		})
	}

}