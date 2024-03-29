package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tesujiro/9gc/ast"
	"github.com/tesujiro/9gc/parser"
	"github.com/tesujiro/9gc/vm"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("引数の個数が正しくありません\n")
	}

	ast.Init() // TODO: env

	src := os.Args[1]
	parser.Tokenize(src)
	parser.Parse()

	// Init
	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global main\n")
	fmt.Printf("main:\n")

	// Prologue: allocate 26 variables in stack
	fmt.Printf("  push rbp\n")
	fmt.Printf("  mov rbp, rsp\n")
	fmt.Printf("  sub rsp, %d\n", (ast.VarCount+1)*8)

	for _, ast := range ast.Code {
		vm.Gen(ast)
		fmt.Printf("  pop rax\n")
	}

	// Epilogue: return last value in RAX
	fmt.Printf("  mov rsp, rbp\n")
	fmt.Printf("  pop rbp\n")
	fmt.Printf("  ret\n")
	return
}
