package parser

import "github.com/tesujiro/9gc/ast"

func Parse() {
	program()
	return
}

func program() {
	for !consume(TK_EOF) {
		ast.Code = append(ast.Code, stmt())
	}
}

func stmt() *ast.Node {
	var node *ast.Node
	if consume(TK_RETURN) {
		node = &ast.Node{
			Ty:  ast.ND_RETURN,
			Lhs: expr(),
		}
	} else {
		node = expr()
	}
	if !consume(';') {
		errorAt(tokens[pos].loc, "the token is not ';'")
	}
	return node
}

func expr() *ast.Node {
	return assign()
}

func assign() *ast.Node {
	node := equality()
	if consume('=') {
		node = ast.NewNode('=', node, assign())
	}
	return node
}

func equality() *ast.Node {
	node := relational()
	for {
		if consume(TK_EQ) {
			node = ast.NewNode(ast.ND_EQ, node, relational())
		} else if consume(TK_NE) {
			node = ast.NewNode(ast.ND_NE, node, relational())
		} else {
			return node
		}
	}
}

func relational() *ast.Node {
	node := add()
	for {
		switch {
		case consume('<'):
			node = ast.NewNode('<', node, add())
		case consume(TK_LE):
			node = ast.NewNode(ast.ND_LE, node, add())
		case consume('>'):
			node = ast.NewNode('>', node, add())
		case consume(TK_GE):
			node = ast.NewNode(ast.ND_GE, node, add())
		default:
			return node
		}
	}
}

func add() *ast.Node {
	//fmt.Println("expr")
	node := mul()
	for {
		if consume('+') {
			node = ast.NewNode('+', node, mul())
		} else if consume('-') {
			node = ast.NewNode('-', node, mul())
		} else {
			return node
		}
	}
}

func mul() *ast.Node {
	//fmt.Println("mul")
	node := unary()
	for {
		if consume('*') {
			node = ast.NewNode('*', node, unary())
		} else if consume('/') {
			node = ast.NewNode('/', node, unary())
		} else {
			return node
		}
	}
}

func unary() *ast.Node {
	if consume('+') {
		return term()
	} else if consume('-') {
		return ast.NewNode('-', ast.NewNodeNum(0), term())
	} else {
		return term()
	}
}

func term() *ast.Node {
	//fmt.Println("term")
	if consume('(') {
		node := expr()
		if !consume(')') {
			errorAt(pos, "no closing parenthesis")
		}
		return node
	}
	if tokens[pos].ty == TK_IDENT {
		varname := tokens[pos].name
		var varNo int

		if v, ok := ast.VarMap.Get(varname); ok {
			varNo = v.(int)
		} else {
			varNo = ast.VarCount
			ast.VarCount++
			ast.VarMap.Put(varname, varNo)
		}

		n := ast.Node{
			Ty:     ast.ND_LVAR,
			Offset: (varNo + 1) * 8,
		}
		pos++
		return &n
	}
	if tokens[pos].ty == TK_NUM {
		val := tokens[pos].val
		pos++
		return ast.NewNodeNum(val)
	}
	errorAt(pos, "not number nor parenthesis")
	return nil
}
