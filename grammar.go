package main

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
