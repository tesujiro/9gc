package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tesujiro/9gc/parser"
	"github.com/tesujiro/9gc/vm"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("引数の個数が正しくありません\n")
	}

	src := os.Args[1]
	parser.Tokenize(src)
	ast := parser.Parse()

	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global main\n")
	fmt.Printf("main:\n")

	vm.Gen(ast)

	fmt.Printf("  pop rax\n")
	fmt.Printf("  ret\n")
	return
}
