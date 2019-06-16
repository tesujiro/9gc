package main

import "fmt"

func expr() *Node {
	//fmt.Println("expr")
	node := mul()
	for {
		if consume('+') {
			node = newNode('+', node, mul())
		} else if consume('-') {
			node = newNode('-', node, mul())
		} else {
			return node
		}
	}
}

func mul() *Node {
	//fmt.Println("mul")
	node := term()
	for {
		if consume('*') {
			node = newNode('*', node, term())
		} else if consume('/') {
			node = newNode('/', node, term())
		} else {
			return node
		}
	}
}

func term() *Node {
	//fmt.Println("term")
	if consume('(') {
		node := expr()
		if !consume(')') {
			errorAt(pos, "no closing parenthesis")
		}
		return node
	}

	if tokens[pos].ty == TK_NUM {
		val := tokens[pos].val
		pos++
		return newNodeNum(val)
	}
	errorAt(pos, "not number nor parenthesis")
	return nil
}

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
