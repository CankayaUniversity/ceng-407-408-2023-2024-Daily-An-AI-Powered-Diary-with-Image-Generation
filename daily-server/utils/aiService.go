package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GetDataFromFlask(daily string) (map[string]interface{}, error) {
	url := "http://localhost:5000"

	payload := map[string]string{
		"daily": daily,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}