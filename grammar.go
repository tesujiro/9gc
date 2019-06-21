package main

func expr() *Node {
	return equality()
}

func equality() *Node {
	node := relational()
	for {
		if consume(TK_EQ) {
			node = newNode(ND_EQ, node, relational())
		} else if consume(TK_NE) {
			node = newNode(ND_NE, node, relational())
		} else {
			return node
		}
	}
}

func relational() *Node {
	node := add()
	for {
		switch {
		case consume('<'):
			node = newNode('<', node, add())
		case consume(TK_LE):
			node = newNode(ND_LE, node, add())
		case consume('>'):
			node = newNode('>', node, add())
		case consume(TK_GE):
			node = newNode(ND_GE, node, add())
		default:
			return node
		}
	}
}

func add() *Node {
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
	node := unary()
	for {
		if consume('*') {
			node = newNode('*', node, unary())
		} else if consume('/') {
			node = newNode('/', node, unary())
		} else {
			return node
		}
	}
}

func unary() *Node {
	if consume('+') {
		return term()
	} else if consume('-') {
		return newNode('-', newNodeNum(0), term())
	} else {
		return term()
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
