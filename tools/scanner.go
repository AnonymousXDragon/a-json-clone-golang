package tools

import (
	"fmt"
	"os"
	reg "regexp"
)

type Scanner struct {
	Source    string
	HadError  bool
	ListToken []Token
	Current   int
	Start     int
	Line      int
}

func NewScanner(source string) Scanner {
	return Scanner{
		Source:    source,
		HadError:  false,
		ListToken: []Token(nil),
		Current:   0,
		Start:     0,
		Line:      0,
	}
}

func (s *Scanner) IsEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *Scanner) ScanTokens() {
	for {
		if !s.IsEnd() {
			s.ScanToken()
		} else {
			break
		}
	}

	s.ListToken = append(s.ListToken, NewToken(EOF, s.Line, "End Of File", s.Current))
}

func (s *Scanner) ScanToken() {
	symbol := s.MoveCursor()
	// fmt.Println("symbol", symbol)

	switch symbol {
	case "{":
		s.AddToken(LeftCurl, "{")
	case "}":
		fmt.Println("...")
		s.AddToken(RightCurl, "}")
	case " ":
		// fmt.Println("space clicked")
		s.AddToken(WHITESPACE, " ")
	case ":":
		// fmt.Println("semi column clicked")
		s.AddToken(SEMICOLUMN, ":")
	case "\"":
		s.AddString()
	case ",":
		s.AddToken(COMMA, ",")
		s.Line++
	case "\n":
		s.Line++
	default:
		if s.IsDigit(symbol) {
			s.NumberProcess(symbol)
		}

		if symbol == "t" || symbol == "f" || symbol == "n" {
			s.DetectKeywords(symbol)
		}
		// else {
		// 	msg := fmt.Sprintf("unexpected character found at %d %s", s.Current, symbol)
		// 	s.HadError = true
		// 	panic(msg)
		// }
	}
}

func (s *Scanner) Peek() string {
	if s.IsEnd() {
		return "\x00"
	}
	return string(s.Source[s.Current])
}

func (s *Scanner) NumberProcess(n string) {
	fmt.Println("number process !!")
	number := n

	for s.IsDigit(s.Peek()) {
		// fmt.Println(number)
		number += s.MoveCursor()
	}

	if s.Peek() == "." && s.IsDigit(s.PeekNext()) {
		number += "."
		s.MoveCursor()
		fmt.Println(string(s.Source[s.Current]))
		for s.IsDigit(s.Peek()) {
			number += s.MoveCursor()
		}
	}

	s.AddToken(Number, number)
	s.MoveCursor()
}

func (s *Scanner) PeekNext() string {
	if s.IsEnd() {
		return "\x00"
	}

	return string(s.Source[s.Current+1])
}

func (s *Scanner) DetectKeywords(start string) {
	text := start
	for s.Peek() != "," && !s.IsEnd() {
		text += s.Peek()
		if s.Peek() == "\n" {
			s.Line++
		}

		if s.IsEnd() {
			msg := fmt.Sprintf("%d unexpected string literal found", s.Current)
			panic(msg)
		}

		s.MoveCursor()
	}

	if text == "true" {
		s.AddToken(True, text)
	} else if text == "null" {
		s.AddToken(Null, text)
	} else {
		s.AddToken(False, text)
	}
	s.Line++
}

func (s *Scanner) IsDigit(c string) bool {
	found, err := reg.Match("[0-9]", []byte(c))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return found
}

func (s *Scanner) AddString() {
	text := ""
	for s.Peek() != "\"" && !s.IsEnd() {
		text += s.Peek()
		if s.Peek() == "\n" || s.Peek() == "," {
			s.Line++
		}

		if s.IsEnd() {
			msg := fmt.Sprintf("%d unexpected string literal found", s.Current)
			s.HadError = true
			panic(msg)
		}
		s.MoveCursor()
	}
	if text == " " || text == "" {
		s.AddToken(EMPTYSTRING, text)
	} else {
		s.AddToken(String, text)
	}
	s.Current++
}

func (s *Scanner) AddToken(token TokenType, text string) {
	s.ListToken = append(s.ListToken, NewToken(token, s.Line, text, s.Current))
}

func (s *Scanner) MoveCursor() string {
	s.Current++
	return string(s.Source[s.Current-1])
}
