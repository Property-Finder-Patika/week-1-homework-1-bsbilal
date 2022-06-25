package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(HasPrefix("selam", "ab"))
	fmt.Println(HasSuffix("Selamms", "ab"))
	fmt.Println(Contains("Selamm", "ab"))

	s := "Hello, BF"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, BF" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	n := 0
	for _, _ = range s {
		n++
	}

	b := 0
	for range s {
		b++
	}

	// "program" in Japanese katakana
	x := ">+=@?"
	fmt.Printf("% x\n", x) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "C"

	fmt.Println(string(1234567)) // "?"

}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	//suffix check
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
func Contains(s, substr string) bool {
	//contains check
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
