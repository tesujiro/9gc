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
	case ND_EQ:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  sete al\n")
		fmt.Printf("  movzb rax,al\n")
	case ND_NE:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setne al\n")
		fmt.Printf("  movzb rax,al\n")
	case '<':
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setl al\n")
		fmt.Printf("  movzb rax,al\n")
	case ND_LE:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setle al\n")
		fmt.Printf("  movzb rax,al\n")
	case '>':
		fmt.Printf("  cmp rdi, rax\n")
		fmt.Printf("  setl al\n")
		fmt.Printf("  movzb rax,al\n")
	case ND_GE:
		fmt.Printf("  cmp rdi, rax\n")
		fmt.Printf("  setle al\n")
		fmt.Printf("  movzb rax,al\n")
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
