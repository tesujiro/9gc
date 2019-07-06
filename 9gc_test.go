package main

import (
	"testing"

	"github.com/tesujiro/9gc/ast"
)

/*
func expect(t *testing.T, expected, actual interface{}, ok bool) {
	if !ok && expected == nil ||
		ok && expected == actual {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	t.Fatalf("%s:%d: %v expected, but got %v\n", file, line, expected, actual)
	return
}
*/

func TestMap(t *testing.T) {
	m := ast.NewMap()
	if _, ok := m.Get("foo"); ok {
		t.Fatalf("Map.Get test error: expected not ok got ok\n")
	}

	m.Put("foo", interface{}(2))
	if actual, ok := m.Get("foo"); !ok {
		t.Fatalf("Map.Get test error: expected ok got not ok\n")
	} else if actual.(int) != 2 {
		t.Fatalf("Map.Get test error: expected 2 got %v\n", actual)
	}
	//fmt.Printf("Get:%v\n", *m.Get("foo"))
}
