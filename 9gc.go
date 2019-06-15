package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("引数の個数が正しくありません\n")
	}

	fmt.Printf(".intel_syntax noprefix\n")
  	fmt.Printf(".global main\n")
  	fmt.Printf("main:\n")
  	expr,err :=strconv.Atoi(os.Args[1])
	if err!=nil{
		log.Fatalf("引数が正しくありません\n")
	}
  	fmt.Printf("  mov rax, %d\n", expr)
  	fmt.Printf("  ret\n")
  	return 
}
