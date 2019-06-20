package main

import (
	"log"
	"strconv"
	"unicode"
)

const (
	TK_NUM = 256 + iota
	TK_EOF
	TK_EQ // ==
	TK_NE // !=
	TK_LE // <=
	TK_GE // >=
)

type Token struct {
	ty    int    // type of token
	val   int    // value when ty is TK_NUM
	input string // token string
	loc   int    // location
}

const (
	ND_NUM = 256 + iota
)

type Node struct {
	ty  int   // type of node
	lhs *Node // left-hand side
	rhs *Node // reft-hand side
	val int   // value when ty is ND_NUM
}

func newNode(ty int, lhs, rhs *Node) *Node {
	return &Node{
		ty:  ty,
		lhs: lhs,
		rhs: rhs,
	}
}

func newNodeNum(val int) *Node {
	//fmt.Printf("newNodeNum(%v)\n", val)
	return &Node{
		ty:  ND_NUM,
		val: val,
	}
}

func consume(ty int) bool {
	if tokens[pos].ty != ty {
		return false
	}
	pos++
	return true
}

var (
	user_input string
	tokens     [100]Token
	pos        int
)

func errorPrint(fmt string, args ...interface{}) {
	log.Fatalf(fmt+"\n", args...)
}

func errorAt(pos int, msg string) {
	log.Printf("%s", user_input)
	space := ""
	for i := 0; i < pos; i++ {
		space += " "
	}
	log.Fatalf("%s^ %s\n", space, msg)
}

func tokenize() {
	i := 0
	p := 0
	opList := []rune{'+', '-', '*', '/', '(', ')'}
loop:
	for p < len(user_input) {
		r := rune(user_input[p])
		if unicode.IsSpace(r) {
			p++
			//fmt.Println("SPACE")
			continue
		}
		for _, ch := range opList {
			if r == ch {
				tokens[i].ty = int(user_input[p])
				tokens[i].input = user_input[p : p+1]
				tokens[i].loc = p
				i++
				p++
				//fmt.Println("PLUS MINUS")
				continue loop
			}
		}
		if unicode.IsDigit(r) {
			tokens[i].ty = TK_NUM
			start := p
			p++
			for p < len(user_input) && unicode.IsDigit(rune(user_input[p])) {
				p++
			}
			stop := p
			digit := user_input[start:stop]

			d, _ := strconv.Atoi(digit)
			tokens[i].input = digit
			tokens[i].val = d
			tokens[i].loc = p
			i++
			//fmt.Println("DIGIT")
			continue
		}
		errorAt(p, "tokenize failed")

		/*
			// Skip line comment
			if user_input[p:p+2] == "//" {
				p += 2
				for user_input[p] != '\n' {
					p++
				}
				continue
			}
			// Skip block omment
			if user_input[p:p+2] == "/*" {
				p += 2
				idx := strings.Index(user_input[p:], "* /")
				if idx < 0 {
					error_at(p, "Comment not terminated")
				}
				p = idx + 2
				continue
			}
		*/
	}
	tokens[i].ty = TK_EOF
	tokens[i].input = "EOF"
	tokens[i].loc = p
	//fmt.Println("EOF :", i)
}
