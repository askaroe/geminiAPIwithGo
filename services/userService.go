package services

import (
	"context"
	"fmt"
	"github.com/askaroe/geminiApiGolang/model"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"os"
	"strings"
)

func (s *Service) SummarizeUserInfo(user *model.User) (string, error) {
	s.Logger.PrintInfo("body", map[string]string{"user": user.String()})

	request := fmt.Sprintf(
		"Give the summarize for the future diet and health of the person, "+
			"with limit of words 50, with age %d, with height %d cm, and with weight %d kg",
		user.Age, user.Height, user.Weight,
	)

	s.Logger.PrintInfo("requestAI", map[string]string{"body": request})

	resp, err := SendUserInfo(&request)
	return resp, err
}

func SendUserInfo(info *string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))

	if err != nil {
		return "", err
	}

	defer client.Close()

	aiModel := client.GenerativeModel("gemini-1.5-flash")

	resp, err := aiModel.GenerateContent(ctx, genai.Text(*info))
	if err != nil {
		return "", nil
	}

	fmt.Println("--> Response")

	var builder strings.Builder

	if resp != nil {
		candidates := resp.Candidates
		if candidates != nil {
			for _, candidate := range candidates {
				content := candidate.Content

				if content != nil {
					text := content.Parts[0]

					fmt.Fprintf(&builder, "%v", text)
				}
			}
		}
	}

	return builder.String(), nil
}
