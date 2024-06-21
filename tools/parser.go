package tools

import (
	"fmt"
	"strconv"
)

/**

object {
	type:"object"
	children: [
		Property{
			type: "Property",
			key: Identifier {
				type: "Identifier"
				value: "key"
			},
			value: Literal {
				type: "Literal",
				value: "keys"
			}
		},
		Property{}
	]
}

**/

type Identifier struct {
	_type string
	value string
}

type Booleans struct {
	_type string
	value bool
}

type Numbers struct {
	_type string
	value float64
}

type Literal struct {
	_type string
	value string
}

type Property struct {
	_type string
	key   interface{}
	value interface{}
}

type Object struct {
	_type     string
	childrens []Property
}

type Parser struct {
	Tokens            []Token
	CurrentTokenIndex int
	Ast               Object
}

func (p *Parser) CheckIsObject() bool {
	var check string

	leftCurl := p.Tokens[0].Literal
	rightCurl := p.Tokens[len(p.Tokens)-2].Literal

	check = fmt.Sprintf("%[1]v", leftCurl+rightCurl)
	return check == "{}"
}

func (p *Parser) Parse() {
	if p.IsEmpty() {
		panic("there are no tokens to proceed !!")
	}

	if !p.CheckIsObject() {
		panic("invalid json format")
	}

	p.Ast._type = "object"

	for !p.IsEnd() {
		p.Match()
	}
}

// func (p *Parser) generatePair() []Token {
// 	var pair []Token
// 	for {
// 	}
// }

// func (p *Parser) Peek() Token {
// 	current := p.tokens[p.currentTokenIndex]
// 	for current.Type != SEMICOLUMN {
// 		p.currentTokenIndex++
// 	}
// 	p.currentTokenIndex++
// }

func (p *Parser) IsEmpty() bool {
	return len(p.Tokens) <= 0
}

func (p *Parser) IsEnd() bool {
	return p.CurrentTokenIndex >= len(p.Tokens)
}

func (p *Parser) Match() {
	current := p.CurrentToken()

	switch current.Type {
	case LeftCurl:
		p.ConsumeToken()
	case RightCurl:
		p.ConsumeToken()
	case WHITESPACE:
		p.ConsumeToken()
	case String:
		p.ConsumeToken()
	case Number:
		p.ConsumeToken()
	case EMPTYSTRING:
		p.ConsumeToken()
	case SEMICOLUMN:
		// var key_type Token
		var value_type interface{}

		key := p.Tokens[p.CurrentTokenIndex-1]
		value := p.Tokens[p.CurrentTokenIndex+1]

		if key.Type != String {
			panic("key type should be a string")
		}

		if value.Type == String || value.Type == WHITESPACE {
			value_type = Literal{
				_type: "Literal",
				value: value.Literal,
			}
		} else if value.Type == True || value.Type == False {
			boolVal, _ := strconv.ParseBool(value.Literal)
			value_type = Booleans{
				_type: "bool",
				value: boolVal,
			}
		} else {
			val, err := strconv.ParseFloat(value.Literal, 64)
			if err != nil {
				panic("number conversion failed")
			}
			value_type = Numbers{
				_type: "Number",
				value: val,
			}
		}

		p.Ast.childrens = append(p.Ast.childrens, Property{
			_type: "Property",
			key: Identifier{
				_type: "Identifier",
				value: key.Literal,
			},
			value: value_type,
		})

		p.ConsumeToken()
	case EOF:
		fmt.Println("...end")
		p.ConsumeToken()
	default:
		panic("parsing failed")
	}

	// if p.CurrentToken().Type == expectedType {
	// 	p.ConsumeToken()
	// } else {
	// 	msg := fmt.Sprintf("SYNTAX ERROR: expected %v but found %v", expectedType, p.CurrentToken().Type)
	// 	panic(msg)
	// }
}

func (p *Parser) IsKeyIdentifier() bool {
	return p.Tokens[p.CurrentTokenIndex+1].Type == SEMICOLUMN
}

func (p *Parser) ConsumeToken() {
	p.CurrentTokenIndex++
}

func (p *Parser) CurrentToken() Token {
	return p.Tokens[p.CurrentTokenIndex]
}
