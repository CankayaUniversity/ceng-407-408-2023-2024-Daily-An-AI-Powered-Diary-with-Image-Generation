package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type TextToImageImage struct {
	Base64       string `json:"base64"`
	Seed         uint32 `json:"seed"`
	FinishReason string `json:"finishReason"`
}

type TextToImageResponse struct {
	Images []TextToImageImage `json:"artifacts"`
}

type TextPrompt struct {
	Text string `json:"text"`
}

type RequestData struct {
	TextPrompts []TextPrompt `json:"text_prompts"`

	CfgScale int `json:"cfg_scale"`
	Height   int `json:"height"`
	Width    int `json:"width"`
	Samples  int `json:"samples"`
	Steps    int `json:"steps"`
}

func createRequestData2(prompt string, height, width int) ([]byte, error) {
	data := RequestData{
		TextPrompts: []TextPrompt{
			{Text: prompt},
		},
		CfgScale: 15,
		Height:   height,
		Width:    width,
		Samples:  1,
		Steps:    40,
	}

	// Encode the data to JSON
	requestData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return requestData, nil
}

func createRequestData(prompt string) []byte {
	return []byte(fmt.Sprintf(`{
		"text_prompts": [
			{
				"text": "%s"
			}
		],
		"cfg_scale": 9,
		"height": 1024,
		"width": 1024,
		"samples": 1,
		"steps":30 
	}`, prompt))
}

func GenerateImage(prompt string) (TextToImageImage, error) {
	engineId := "stable-diffusion-xl-1024-v1-0"
	apiHost := "https://api.stability.ai"
	reqUrl := apiHost + "/v1/generation/" + engineId + "/text-to-image"
	apiKey, hasAPIKey := os.LookupEnv("STABILITY_API_KEY")
	if !hasAPIKey {
		return TextToImageImage{}, errors.New("Missing STABILITY_API_KEY environment variable")
	}

	logger := logrus.New()
	logger.Infof(apiKey)

	data := createRequestData(prompt)
	logger.Infof("request parameters %v", data)
	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	// Execute the request & read all the bytes of the body
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	logger.Infof("%v\n%v", res, err)

	if res.StatusCode != 200 {
		return TextToImageImage{}, errors.New("response not 200")
	}

	// Decode the JSON body
	var body TextToImageResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return TextToImageImage{}, errors.New("couldn't decode json body")

	}

	return body.Images[0], nil
}
