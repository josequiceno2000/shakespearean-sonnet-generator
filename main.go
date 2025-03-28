package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
) 


func main() {
	// Get dictionary
	cmuDict, err := loadCMUDictionary()
	if err != nil {
		fmt.Println("Error loading dictionary:", err)
		os.Exit(1)
	}

	// Get number of sentences to analyze
	if len(os.Args) < 2 {
		fmt.Println("Usage: ssg <name>")
		os.Exit(1)
	}
	arg1 := os.Args[1]
	maxLength, err := strconv.Atoi(arg1)
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid integer.\n", arg1)
		os.Exit(1)
	}

	// Set up the sonnet
	sonnet := `When I do count the clock that tells the time,\n 
				And see the brave day sunk in hideous night;\n 
				When I behold the violet past prime,\n 
				And sable curls all silver’d o’er with white;\n
				When lofty trees I see barren of leaves\n 
				Which erst from heat did canopy the herd,\n
				And summer’s green all girded up in sheaves\n 
				Borne on the bier with white and bristly beard,\n 
				Then of thy beauty do I question make,\n 
				That thou among the wastes of time must go,\n
				Since sweets and beauties do themselves forsake\n
				And die as fast as they see others grow;\n 
				And nothing ‘gainst Time’s scythe can make defence\n
				Save breed, to brave him when he takes thee hence.`
	
	sonnetLines := strings.Split(sonnet, "\n")
	
	// Take it line by line
	for i := 0; i < maxLength && i < len(sonnetLines); i++ {
		totalSyllables := 0
		line := strings.ReplaceAll(sonnetLines[i], "\r", "")
		line = strings.ReplaceAll(line, "\n", "")
		words := strings.Fields(line)
		for _, word := range(words) {
			cleanWord := strings.TrimFunc(word, func(r rune) bool {
				return !unicode.IsLetter(r) && r != '\''
			})
	
			pronunciation, found := cmuDict[strings.ToUpper(cleanWord)]
			if found {
				syllableCount := countSyllables(pronunciation)
				totalSyllables += syllableCount
				fmt.Printf("%s: %d syllables\n", word, syllableCount)
			} else {
				fmt.Printf("%s: Pronunciation not found\n", cleanWord)
			}
		}

		fmt.Printf("Line %v has a total of %v syllables\n", i + 1, totalSyllables)
	}
}