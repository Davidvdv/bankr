package classification

import (
	"bankr/internal/model"
	"context"
	"fmt"
	"google.golang.org/genai"
	"log"
	"os"
	"strings"
)

type Classifier interface {
	Classify(descriptions []string)
}

type TransactionClassifier struct {
}

func (t *TransactionClassifier) Classify(descriptions []string) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash-lite",
		genai.Text("Could you classify this list of transactions: "+strings.Join(descriptions, ", ")+""+
			" and use the following classification labels: "+strings.Join(classificationsToStrings(model.AllClassifications), ", ")),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())
}

// Helper function to convert []Classification to []string
func classificationsToStrings(classifications []model.Classification) []string {
	result := make([]string, len(classifications))
	for i, c := range classifications {
		result[i] = c.String()
	}
	return result
}
