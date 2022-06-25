package main

import "fmt"

func main() {
	fmt.Print(isAnagram("asdfsa", "dafdaf"))
}

func isAnagram(wordA, wordB string) bool {
	//check words are anagram or not
	if len(wordA) != len(wordB) {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < len(wordA); i++ {
		// iterates and push letters to map
		anagramMap[string(wordA[i])]++
	}

	for i := 0; i < len(wordB); i++ {
		// iterates and remove letters  map

		anagramMap[string(wordB[i])]--
	}

	for i := 0; i < len(wordA); i++ {
		// iterates and check did every word remove from map?
		if anagramMap[string(wordA[i])] != 0 {
			return false
		}
	}
	return true

}
