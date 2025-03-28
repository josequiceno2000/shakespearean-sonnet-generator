package main

import (
	"bufio"
	"os"
	"strings"
)

func loadCMUDictionary() (map[string]string, error) {
	dict := make(map[string]string)
	file, err := os.Open("cmudict-0.7b")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip comments
		if strings.HasPrefix(line, ";;;") {
			continue
		}
		parts := strings.SplitN(line, "  ", 2)
		if len(parts) == 2 {
			dict[parts[0]] = parts[1]
		}
	}
	return dict, scanner.Err()
}