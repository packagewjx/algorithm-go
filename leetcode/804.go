package leetcode

var MORSE = [...]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	transformations := make(map[string]bool)
	count := 0

	for _, word := range words {
		morse := toMorse(word)
		_, ok := transformations[morse]
		if !ok {
			count++
			transformations[morse] = true
		}
	}

	return count
}

func toMorse(word string) string {
	morse := ""
	for _, c := range word {
		morse += MORSE[c-'a']
	}
	return morse
}
