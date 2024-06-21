package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"jparse/tools"
	"jparse/utils"
	//	"strings"
)

var (
	println = fmt.Println
	printf  = fmt.Printf
)

func main() {
	file := os.Args[1]
	data, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}

	var source string
	scaner := bufio.NewScanner(data)
	scaner.Split(bufio.ScanLines)
	// var LexerExe tools.Scanner

	for scaner.Scan() {
		source += scaner.Text() + "\n"
	}

	if len(source) == 0 {
		fmt.Println("invalid json")
		os.Exit(1)
	}
	source = strings.TrimSpace(source)
	println("total length: ", len(source))
	// println(source)
	// println("source ", source)
	// println("checking...", string(source[1:5]))
	LexerExe := tools.NewScanner(source)
	LexerExe.ScanTokens()
	// isLastRightCurl := LexerExe.ListToken[len(LexerExe.ListToken)-2]
	// println("Right Curl", isLastRightCurl)
	// println("key:", LexerExe.ListToken[2-1])
	// println("semi-column", LexerExe.ListToken[2])
	// println("value:", LexerExe.ListToken[2+1])
	isLastComma := LexerExe.ListToken[len(LexerExe.ListToken)-3].Type == tools.COMMA
	if isLastComma {
		fmt.Println("invalid json , remove the last comma ")
		os.Exit(1)
	}

	lTokens := utils.Filter(LexerExe.ListToken, func(t tools.TokenType) bool {
		return t == tools.WHITESPACE
	})
	// printf("%#v \n", LexerExe.ListToken)
	printf("improved %#v \n", lTokens)
	// parsing

	parser := tools.Parser{
		Tokens:            lTokens,
		CurrentTokenIndex: 0,
	}

	parser.Parse()
	fmt.Printf("tree %#v", parser.Ast)

	defer data.Close()
}
