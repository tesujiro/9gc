package vm

import (
	"fmt"
	"log"

	"github.com/tesujiro/9gc/ast"
)

func errorPrint(fmt string, args ...interface{}) {
	log.Fatalf(fmt+"\n", args...)
}

func genLval(node *ast.Node) {
	if node.Ty != ast.ND_LVAR {
		errorPrint("Left value is not a local variable.")
	}
	fmt.Printf("  mov rax, rbp\n")
	fmt.Printf("  sub rax, %d\n", node.Offset)

	fmt.Printf("  push rax\n")
}

func Gen(node *ast.Node) {
	switch node.Ty {
	case ast.ND_NUM:
		fmt.Printf("  push %d\n", node.Val)
		return
	case ast.ND_LVAR:
		genLval(node)
		fmt.Printf("  pop rax\n")
		fmt.Printf("  mov rax, [rax]\n")
		fmt.Printf("  push rax\n")
		return
	case '=':
		genLval(node.Lhs)
		Gen(node.Rhs)
		fmt.Printf("  pop rdi\n")
		fmt.Printf("  pop rax\n")
		fmt.Printf("  mov [rax], rdi\n")
		fmt.Printf("  push rdi\n")
		return
	case ast.ND_RETURN:
		Gen(node.Lhs)
		fmt.Printf("  pop rax\n")
		fmt.Printf("  mov rsp, rbp\n")
		fmt.Printf("  pop rbp\n")
		fmt.Printf("  ret\n")
		return
	}
	Gen(node.Lhs)
	Gen(node.Rhs)
	fmt.Printf("  pop rdi\n")
	fmt.Printf("  pop rax\n")
	switch node.Ty {
	case ast.ND_EQ:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  sete al\n")
		fmt.Printf("  movzb rax,al\n")
	case ast.ND_NE:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setne al\n")
		fmt.Printf("  movzb rax,al\n")
	case '<':
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setl al\n")
		fmt.Printf("  movzb rax,al\n")
	case ast.ND_LE:
		fmt.Printf("  cmp rax, rdi\n")
		fmt.Printf("  setle al\n")
		fmt.Printf("  movzb rax,al\n")
	case '>':
		fmt.Printf("  cmp rdi, rax\n")
		fmt.Printf("  setl al\n")
		fmt.Printf("  movzb rax,al\n")
	case ast.ND_GE:
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
