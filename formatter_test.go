package formatter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWarning(t *testing.T) {
	result := Warning("test")
	if result != "\u001B[31mtest\u001B[0m" {
		t.Error("incorrect output:", []byte(result))
	}
	result2 := Warning("")
	if result2 != "\u001B[31m\u001B[0m" {
		t.Error("incorrect output:", []byte(result2))
	}
}

func TestSpace(t *testing.T) {
	data := []struct {
		name   string
		length int
		in     []string
		out    string
	}{
		{"zero", 1, nil, ""},
		{"one", 1, []string{"hello"}, "hello"},
		{"two", 20, []string{"hello", "jon"}, "hello            jon"},
		{"three", 20, []string{"hello", "there", "jon"}, "hello    there   jon"},
		{"extra", 10, []string{"hello", "there", "jon"}, "hello there jon"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			out := Space(d.length, d.in...)
			if diff := cmp.Diff(out, d.out); diff != "" {
				t.Error(diff)
			}
		})
	}
}
