package utils

import (
	"jparse/tools"
)

func Filter(elems []tools.Token, f func(ind tools.TokenType) bool) []tools.Token {
	arr := make([]tools.Token, len(elems))
	copy(arr, elems)

	// fmt.Println("before", arr)
	for i := 0; i < len(arr); {
		if f(arr[i].Type) {
			arr = append(arr[:i], arr[i+1:]...)
			i = 0
		}
		i++
	}

	// fmt.Println(arr)
	return arr
}
