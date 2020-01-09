package test

import "testing"

func hello() string{
	return "hello world"
}

func Test_Standard(t *testing.T) {
	 got:=hello()
	 want:="hello go"
	if got !=want {
		t.Errorf("got '%q' want '%q'",got,want)
	}
}
