package slack

import (
	"io"
	"net/http"
)

func SendMessage(url string, message io.Reader) error {
	_, err := http.Post(
		url,
		"application/json",
		message,
	)
	if err != nil {
		return err
	}

	return nil
}
