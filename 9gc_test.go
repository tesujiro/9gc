package main

import (
	"runtime"
	"testing"

	"github.com/tesujiro/9gc/parser"
)

func expect(t *testing.T, expected, actual *interface{}) {
	if expected == actual && expected == nil ||
		expected != nil && actual != nil && *expected == *actual {
		return
	}
	_, file, line, _ := runtime.Caller(1)
	t.Fatalf("%s:%d: %v expected, but got %v\n", file, line, *expected, *actual)
	return
}

func TestMap(t *testing.T) {
	m := parser.NewMap()
	expect(t, nil, m.Get("foo"))
	//fmt.Printf("Get:%v\n", m.Get("foo"))

	m.Put("foo", interface{}(2))
	ex := interface{}(2)
	expect(t, &ex, m.Get("foo"))
	//fmt.Printf("Get:%v\n", *m.Get("foo"))
}
