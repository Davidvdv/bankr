package classification

import (
	"regexp"
	"strings"
)

type Classification string

const (
	CategoryFood          Classification = "Food"
	CategoryUtilities     Classification = "Utilities"
	CategoryHobbies       Classification = "Hobbies"
	CategoryCar           Classification = "Car"
	CategoryShopping      Classification = "Shopping"
	CategoryHealth        Classification = "Health"
	CategoryEntertainment Classification = "Entertainment"
	CategoryTransport     Classification = "Transport"
	CategoryOther         Classification = "Other"
)

type TransactionDto struct {
	Description string
	Amount      float64
	Category    Classification
	Confidence  float64 // 0.0 to 1.0 indicating how confident the classification is
}

type ClassificationContext struct {
	patterns map[Classification][]string
	keywords map[Classification][]string
}

func NewCategorizer() *ClassificationContext {
	return &ClassificationContext{
		patterns: map[Classification][]string{
			CategoryFood: {
				`(?i)(restaurant|cafe|coffee|pizza|burger|mcdonalds|kfc|subway|starbucks|dominos)`,
				`(?i)(grocery|supermarket|walmart|target|safeway|kroger|food|dining)`,
				`(?i)(bar|pub|brewery|wine|liquor)`,
			},
			CategoryUtilities: {
				`(?i)(electric|electricity|gas|water|internet|phone|cable|utility)`,
				`(?i)(verizon|att|comcast|spectrum|pg&e|edison)`,
				`(?i)(heating|cooling|sewer|trash|garbage)`,
			},
			CategoryCar: {
				`(?i)(gas|gasoline|fuel|shell|chevron|exxon|bp|mobil)`,
				`(?i)(car|auto|mechanic|repair|service|oil change|tire)`,
				`(?i)(insurance|registration|dmv|parking|toll)`,
			},
			CategoryHobbies: {
				`(?i)(hobby|craft|art|supplies|music|instrument|book|games)`,
				`(?i)(amazon|ebay|etsy|hobby lobby|michaels)`,
				`(?i)(photography|camera|sports|equipment|gym|fitness)`,
			},
			CategoryShopping: {
				`(?i)(shopping|retail|store|mall|outlet|boutique)`,
				`(?i)(clothing|shoes|fashion|apparel|department)`,
				`(?i)(home|furniture|decoration|garden|hardware)`,
			},
			CategoryHealth: {
				`(?i)(doctor|medical|hospital|pharmacy|cvs|walgreens)`,
				`(?i)(dental|dentist|health|medicine|prescription|clinic)`,
				`(?i)(insurance|copay|deductible)`,
			},
			CategoryEntertainment: {
				`(?i)(movie|cinema|theater|netflix|spotify|streaming)`,
				`(?i)(concert|show|event|ticket|entertainment)`,
				`(?i)(game|gaming|xbox|playstation|nintendo)`,
			},
			CategoryTransport: {
				`(?i)(uber|lyft|taxi|bus|train|subway|transit)`,
				`(?i)(airline|flight|airport|travel|hotel)`,
				`(?i)(rental|hertz|enterprise|avis)`,
			},
		},
		keywords: map[Classification][]string{
			CategoryFood:          {"food", "eat", "restaurant", "cafe", "grocery", "dining", "meal"},
			CategoryUtilities:     {"electric", "gas", "water", "internet", "phone", "utility", "bill"},
			CategoryCar:           {"gas", "fuel", "car", "auto", "parking", "repair", "insurance"},
			CategoryHobbies:       {"hobby", "craft", "art", "music", "book", "game", "sport", "gym"},
			CategoryShopping:      {"shop", "store", "retail", "buy", "purchase", "clothing", "shoes"},
			CategoryHealth:        {"doctor", "medical", "pharmacy", "health", "dental", "medicine"},
			CategoryEntertainment: {"movie", "show", "concert", "game", "entertainment", "streaming"},
			CategoryTransport:     {"transport", "travel", "uber", "taxi", "flight", "bus", "train"},
		},
	}
}

func (c *ClassificationContext) Classify(description string) (Classification, float64) {
	cleanDesc := strings.ToLower(strings.TrimSpace(description))

	// Try pattern matching first (higher confidence)
	for category, patterns := range c.patterns {
		for _, pattern := range patterns {
			if matched, _ := regexp.MatchString(pattern, cleanDesc); matched {
				return category, 0.9
			}
		}
	}

	// Try keyword matching (medium confidence)
	bestMatch := CategoryOther
	bestScore := 0.0

	for category, keywords := range c.keywords {
		score := c.calculateKeywordScore(cleanDesc, keywords)
		if score > bestScore {
			bestScore = score
			bestMatch = category
		}
	}

	if bestScore > 0.3 {
		return bestMatch, 0.6
	}

	return CategoryOther, 0.1
}

func (c *ClassificationContext) calculateKeywordScore(description string, keywords []string) float64 {
	words := strings.Fields(description)
	matches := 0

	for _, word := range words {
		cleanWord := cleanWord(word)
		for _, keyword := range keywords {
			if strings.Contains(cleanWord, keyword) || strings.Contains(keyword, cleanWord) {
				matches++
				break
			}
		}
	}

	if len(words) == 0 {
		return 0.0
	}

	return float64(matches) / float64(len(words))
}

func (c *ClassificationContext) ClassifyTransactions(descriptions []string, amounts []float64) []TransactionDto {
	transactions := make([]TransactionDto, len(descriptions))

	for i, desc := range descriptions {
		category, confidence := c.Classify(desc)
		amount := 0.0
		if i < len(amounts) {
			amount = amounts[i]
		}

		transactions[i] = TransactionDto{
			Description: desc,
			Amount:      amount,
			Category:    category,
			Confidence:  confidence,
		}
	}

	return transactions
}

// AddCustomRule allows adding custom categorization rules
func (c *ClassificationContext) AddCustomRule(category Classification, pattern string) {
	if c.patterns[category] == nil {
		c.patterns[category] = []string{}
	}
	c.patterns[category] = append(c.patterns[category], pattern)
}

// GetCategoryStats returns statistics about categorized transactions
func GetCategoryStats(transactions []TransactionDto) map[Classification]ClassificationStats {
	stats := make(map[Classification]ClassificationStats)

	for _, transaction := range transactions {
		stat := stats[transaction.Category]
		stat.Count++
		stat.TotalAmount += transaction.Amount
		stat.AvgConfidence += transaction.Confidence
		stats[transaction.Category] = stat
	}

	// Calculate averages
	for category, stat := range stats {
		if stat.Count > 0 {
			stat.AvgAmount = stat.TotalAmount / float64(stat.Count)
			stat.AvgConfidence = stat.AvgConfidence / float64(stat.Count)
			stats[category] = stat
		}
	}

	return stats
}

// FindLowConfidenceTransactions Helper function to find low-confidence categorizations that might need manual review
func FindLowConfidenceTransactions(transactions []TransactionDto, threshold float64) []TransactionDto {
	var lowConfidence []TransactionDto

	for _, transaction := range transactions {
		if transaction.Confidence < threshold {
			lowConfidence = append(lowConfidence, transaction)
		}
	}

	return lowConfidence
}

// GetTransactionsByCategory Helper function to get transactions by category
func GetTransactionsByCategory(transactions []TransactionDto, category Classification) []TransactionDto {
	var filtered []TransactionDto

	for _, transaction := range transactions {
		if transaction.Category == category {
			filtered = append(filtered, transaction)
		}
	}

	return filtered
}

type ClassificationStats struct {
	Count         int
	TotalAmount   float64
	AvgAmount     float64
	AvgConfidence float64
}
