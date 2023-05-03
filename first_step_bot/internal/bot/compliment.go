package tgbot

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type InputCompliment struct {
	Text string `json:"compliment"`
}

func complimentApi(logger *logrus.Logger) string {
	resp, err := http.Get("https://complimentr.com/api")
	if err != nil {
		return "Internal Error in compliment API"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Internal Error in compliment API"
	}

	result := &InputCompliment{}

	json.Unmarshal(body, result)
	logger.Info(result.Text)

	return result.Text
}
