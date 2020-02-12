func uniqueMorseRepresentations(words []string) int {
	counter := 0
	alphabetToMorseCodeTable := make(map[string]string)
	morseCodelist := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
	}

	for l := 'a'; l <= 'z'; l++ {
		alphabetToMorseCodeTable[string(l)] = morseCodelist[counter]
		counter++
	}

	morseCodeCount := make(map[string]int)

	for _, word := range words {
		var morseCode string
		letters := strings.Split(word, "")
		for _, l := range letters {
			morseCode += alphabetToMorseCodeTable[l]
		}
		morseCodeCount[morseCode] += 1
	}

	return len(morseCodeCount)
}
