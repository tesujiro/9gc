package vm

import (
	"fmt"

	"github.com/tesujiro/9gc/parser"
)

func Gen(node *parser.Node) {
	if node.Ty == parser.ND_NUM {
		fmt.Printf("  push %d\n", node.Val)
		return
	}
	Gen(node.Lhs)
	Gen(node.Rhs)
	fmt.Printf("  pop rdi\n")
	fmt.Printf("  pop rax\n")
	switch node.Ty {
	case parser.ND_EQ:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  sete al\n")
		fmt.Printf("  movzb rax,al\n")
	case parser.ND_NE:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setne al\n")
		fmt.Printf("  movzb rax,al\n")
	case '<':
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setl al\n")
		fmt.Printf("  movzb rax,al\n")
	case parser.ND_LE:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setle al\n")
		fmt.Printf("  movzb rax,al\n")
	case '>':
		fmt.Printf("  cmp rdi, rax\n")
		fmt.Printf("  setl al\n")
		fmt.Printf("  movzb rax,al\n")
	case parser.ND_GE:
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
