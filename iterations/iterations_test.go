package iterations

import (
	"fmt"
	"testing"
)


func TestRepeated(t *testing.T){
	repeated := Repeated("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}



func ExampleRepeated(){
	s := Repeated("b")
   fmt.Println(s)
   // Output: bbbbb
}


func BenchmarkRepeated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeated("a")
	}
}