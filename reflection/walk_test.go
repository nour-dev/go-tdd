package reflection_test

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Alnoor"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d, expected %d", len(got), 1)
	}
	if got[0] != expected {
		t.Errorf("got %q, expected %q", got[0], expected)
	}
}

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
