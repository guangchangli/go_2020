package test

import "testing"

func hello() string {
	return "hello world"
}

func Test_Standard(t *testing.T) {
	got := hello()
	want := "hello go"
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

const str1="hello"
const str2="world"
func Test_Common(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("compare two string", func(t *testing.T) {
		got := "hello"
		want := "hell"
		assertCorrectMessage(t, got, want)
	})
	t.Run("compare use const", func(t *testing.T) {
		assertCorrectMessage(t,str1,str2)
	})
}
