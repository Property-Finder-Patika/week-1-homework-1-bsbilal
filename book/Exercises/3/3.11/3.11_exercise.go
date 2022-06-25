package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("selammsmsms"))
}

func comma(s string) string {
	// find float index

	floatIndex := strings.Index(s, ".")
	if floatIndex == -1 {
		floatIndex = len(s)
	}

	var buf bytes.Buffer

	if strings.HasPrefix(s, "+") ||
		strings.HasPrefix(s, "-") {
		//check number has positive or negative sign
		buf.WriteByte(s[0])
		s = s[1:]
		floatIndex--
	}

	// find mod of float point
	i := floatIndex % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])

	for j, c := range s[i:floatIndex] {
		if j%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(c)
	}
	buf.WriteString(s[floatIndex:])
	return buf.String()
}
