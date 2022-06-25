package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("selammsmsms"))
}

func comma(s string) string {
	// that function splits every three letters
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	lenA := len(s) / 3
	lenB := len(s) % 3
	if lenA != 0 {
		buf.WriteString(s[:lenA])
		buf.WriteByte(',')
	}

	for i := 0; i < lenB; i++ {
		buf.WriteString(s[i*3+lenA : lenA+i*3+3])
		if i != lenB-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
