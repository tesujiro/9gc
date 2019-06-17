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
