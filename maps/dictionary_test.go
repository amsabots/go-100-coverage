package maps

import "testing"

func TestSearch(t *testing.T){
	dictionary:= Dictionary{ "test": "test one"}
	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "test one"
		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatalf("expected an error but got none")
		}
		assertError(t, err, ErrKeyNotFoundInMap)
	})
}


func TestAdd(t *testing.T){
	dict := Dictionary{}
	dict.Add("add", "word")
	want := "word"
	got, err := dict.Search("add")
	if err != nil {
		t.Fatal("expected not to return an error")
	}
	assertStrings(t, got, want)
}


func assertStrings (t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}

func assertError(t testing.TB, got, want error){
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}