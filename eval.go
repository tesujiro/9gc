package main

import "fmt"

func gen(node *Node) {
	if node.ty == ND_NUM {
		fmt.Printf("  push %d\n", node.val)
		return
	}
	gen(node.lhs)
	gen(node.rhs)
	fmt.Printf("  pop rdi\n")
	fmt.Printf("  pop rax\n")
	switch node.ty {
	case '+':
		fmt.Printf("  add rax, rdi\n")
	case '-':
		fmt.Printf("  sub rax, rdi\n")
	case '*':
		fmt.Printf("  imul rax, rdi\n")
	case '/':
		fmt.Printf("  cqo\n")
		fmt.Printf("  idiv rdi\n")
	}
	fmt.Printf("  push rax\n")
}
