package utils

import (
	"encoding/base64"
)

func ImageToBase64(imageData []byte) string {
	encodedImage := "data:image/png;base64," + base64.StdEncoding.EncodeToString(imageData)
	return encodedImage
}
