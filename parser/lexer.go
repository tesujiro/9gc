package parser

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

var (
	user_input string
	tokens     []Token
	pos        int
)

func initTokens() {
	tokens = make([]Token, 0)
}

func pushTokens(t Token) {
	tokens = append(tokens, t)
}

func consume(ty int) bool {
	if tokens[pos].ty != ty {
		return false
	}
	pos++
	return true
}

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

func Tokenize(src string) {
	user_input = src
	initTokens()
	p := 0 // position in src
	opList := []rune{'<', '>', '+', '-', '*', '/', '(', ')'}
	multiCharOpList := []struct {
		ope string
		tok int
	}{
		{"==", TK_EQ},
		{"!=", TK_NE},
		{"<=", TK_LE},
		{">=", TK_GE},
	}

loop:
	for p < len(user_input) {
		r := rune(user_input[p])
		if unicode.IsSpace(r) {
			p++
			//fmt.Println("SPACE")
			continue
		}
		// Tokenize Multi Character Operators
		for _, op := range multiCharOpList {
			if p+len(op.ope) > len(user_input) {
				continue
			}
			s := user_input[p : p+len(op.ope)]
			if s == op.ope {
				pushTokens(Token{
					ty:    op.tok,
					input: s,
					loc:   p,
				})
				p += len(op.ope)
				//fmt.Println("MULIT CHAR OP")
				continue loop
			}
		}
		// Tokenize Single Character Operators
		for _, ch := range opList {
			if r == ch {
				pushTokens(Token{
					ty:    int(user_input[p]),
					input: user_input[p : p+1],
					loc:   p,
				})
				p++
				//fmt.Println("SINGLE CHAR OP")
				continue loop
			}
		}
		// Tokenize Numbers
		if unicode.IsDigit(r) {
			start := p
			p++
			for p < len(user_input) && unicode.IsDigit(rune(user_input[p])) {
				p++
			}
			stop := p
			digit := user_input[start:stop]

			d, _ := strconv.Atoi(digit)
			pushTokens(Token{
				ty:    TK_NUM,
				input: digit,
				val:   d,
				loc:   p,
			})
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
	pushTokens(Token{
		ty:    TK_EOF,
		input: "EOF",
		loc:   p,
	})
	//fmt.Println("EOF :")
}
