package parser

const (
	ND_NUM = 256 + iota
	ND_EQ
	ND_NE
	ND_LE
	ND_GE
)

type Node struct {
	Ty  int   // type of node
	Lhs *Node // left-hand side
	Rhs *Node // reft-hand side
	Val int   // value when ty is ND_NUM
}

func newNode(ty int, lhs, rhs *Node) *Node {
	return &Node{
		Ty:  ty,
		Lhs: lhs,
		Rhs: rhs,
	}
}

func newNodeNum(val int) *Node {
	//fmt.Printf("newNodeNum(%v)\n", val)
	return &Node{
		Ty:  ND_NUM,
		Val: val,
	}
}

func Parse() *Node {
	return expr()
}

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
