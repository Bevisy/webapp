package myatoi

import (
	"log"
)

func myatoi(s string) int {
	s0 := s
	if s[0] == '-' || s[0] == '+' { //"" 与 '' 区别
		s = s[1:]
		if len(s) < 1 {
			log.Fatal("the input is wrong.")
		}
	}

	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			log.Fatal("the input is invalid.")
		}
		n = n*10 + int(ch)
	}

	if s0[0] == '-' {
		n = -n
	}

	return n
}
