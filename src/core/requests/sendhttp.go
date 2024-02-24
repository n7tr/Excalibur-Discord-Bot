package requests

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Sendhttp(url string, method string, data []byte) string {
	godotenv.Load()
	TOKEN := os.Getenv("BOT_TOKEN")
	client := &http.Client{}

	var req *http.Request

	if data != nil {
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(data))
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}

	req.Header.Set("Authorization", "Bot "+TOKEN)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	data, _ = io.ReadAll(resp.Body)

	return string(data)
}
