package utils

import (
	"encoding/json"
	"fmt"
	"github.com/Final-Projectors/daily-server/model"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
)

type PredictionRequest struct {
	Instances []Text `json:"instances"`
}

type Text struct {
	Text string `json:"text"`
}

var (
	logger = logrus.New()
)

func createRequestData(param string) []byte {
	return []byte(fmt.Sprintf(`{
		"instances": [
			{
				"text": "%s"
			}
		]
	}`, param))
}

func GetAIPrediction(text string) (model.Prediction, error) {
	endpointURL := os.Getenv("AI_ENDPOINT")

	logger.Infof("Inside GCP AI function")
	escapedText := strings.ReplaceAll(text, `"`, `\"`)

	authToken := exec.Command("gcloud", "auth", "print-access-token")
	output, err := authToken.Output()
	outstr := string(output[:])
	logger.Infof("%v", outstr)

	escapedOut := strings.ReplaceAll(outstr, `"`, `\"`)

	dataPart := fmt.Sprintf(`{"instances": [{ "text": "%s" }]}`, escapedText)
	auth := fmt.Sprintf("Authorization: Bearer %v", strings.TrimSpace(escapedOut))

	logger.Infof(endpointURL)

	cmd := exec.Command("curl",
		"-X", "POST",
		"-H", auth,
		"-H", "Content-Type: application/json",
		endpointURL,
		"-d", dataPart,
	)

	output, err = cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return model.Prediction{}, nil
	}
	fmt.Println(string(output))

	var predictionResponse model.PredictionResponse
	err = json.Unmarshal([]byte(output), &predictionResponse)
	if err != nil {
		logger.Infof("Error: %v", err)
	}

	return predictionResponse.Predictions[0], nil
}
