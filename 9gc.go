package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("引数の個数が正しくありません\n")
	}

	user_input = os.Args[1]
	tokenize()

	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global main\n")
	fmt.Printf("main:\n")

	if tokens[0].ty != TK_NUM {
		errorAt(tokens[0].loc, "argument is not a number")
	}
	fmt.Printf("  mov rax, %d\n", tokens[0].val)

	i := 1
	for tokens[i].ty != TK_EOF {
		//fmt.Println("tokens[",i,"]:", tokens[i])
		switch tokens[i].ty {
		case '+':
			i++
			if tokens[i].ty != TK_NUM {
				errorAt(tokens[i].loc, "argument after + is not a number")
			}
			fmt.Printf("  add rax, %d\n", tokens[i].val)
			i++
		case '-':
			i++
			if tokens[i].ty != TK_NUM {
				errorAt(tokens[i].loc, "argument after - is not a number")
			}
			fmt.Printf("  sub rax, %d\n", tokens[i].val)
			i++
		default:
			errorAt(tokens[i].loc, "not expected token")
		}

	}

	fmt.Printf("  ret\n")
	return
}
