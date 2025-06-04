package classification

type KnownClassification struct {
	Name     string
	Keywords []string
}

func CreateKnownClassifications() []KnownClassification {
	return []KnownClassification{
		{
			Name:     "Food",
			Keywords: []string{""},
		},
		{
			Name:     "Takeaways",
			Keywords: []string{"takeaways"},
		},
	}
}
