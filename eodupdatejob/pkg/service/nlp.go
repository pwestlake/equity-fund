package service

import (
	"context"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

// NLPService ...
// Natural Language Processing service
type NLPService struct {
	ctx context.Context
}

// NewNLPService ...
// Create function for a NLPService
func NewNLPService() NLPService {
	ctx := context.Background()

	return NLPService{ctx: ctx}
}

// GetSentiment ...
func (s *NLPService) GetSentiment(text *string) (float32, error) {
	client, err := language.NewClient(s.ctx)
	if err != nil {
		return 0, err
	}

	// Limit text to 1000 characters
	size := len(*text)
	if size > 1000 {
		size = 1000
	}

	slice := (*text)[0:size]

	sentiment, err := client.AnalyzeSentiment(s.ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: string(slice),
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		return 0, err
	}

	return sentiment.DocumentSentiment.Score, nil
}
