package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T){
	buf := bytes.Buffer{}
	Greet(&buf, "Andrew")

	got := buf.String()
	want := "Hello, Andrew"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}