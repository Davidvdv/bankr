package classification

import (
	"regexp"
	"strings"
)

type AnalysisStats struct {
	TotalDescriptions int
	AvgLength         float64
	AvgWordCount      float64
	CommonWords       map[string]int
	LengthDistrib     []int
}

func AnalyzeDescriptions(descriptions []string) AnalysisStats {
	stats := AnalysisStats{
		TotalDescriptions: len(descriptions),
		CommonWords:       make(map[string]int),
	}

	totalLength := 0
	totalWords := 0

	for _, desc := range descriptions {
		length := len(desc)
		words := strings.Fields(strings.ToLower(desc))
		wordCount := len(words)

		totalLength += length
		totalWords += wordCount

		// Count word frequency
		for _, word := range words {
			cleanWord := cleanWord(word)
			if len(cleanWord) > 1 {
				stats.CommonWords[cleanWord]++
			}
		}
	}

	stats.AvgLength = float64(totalLength) / float64(len(descriptions))
	stats.AvgWordCount = float64(totalWords) / float64(len(descriptions))

	return stats
}

func cleanWord(word string) string {
	// Remove punctuation and convert to lowercase
	reg := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return reg.ReplaceAllString(strings.ToLower(word), "")
}
