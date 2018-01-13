package service

import (
	"encoding/json"
	"net/http"

	"bytes"
	"fmt"

	"github.com/pkg/errors"
)

type SlackAPIService struct {
	url string
}
type msgModel struct {
	Text string `json:"text"`
}

func NewSlackAPIService(url string) *SlackAPIService {
	return &SlackAPIService{url}
}
func (s *SlackAPIService) SendMsg(msg string) error {
	msgJson, err := json.Marshal(msgModel{msg})
	if err != nil {
		return err
	}
	return s.send(msgJson)

}
func (s *SlackAPIService) send(body []byte) error {
	req, _ := http.NewRequest("POST", s.url, bytes.NewReader(body))
	req.Header.Set("Content-type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println()
	if resp.StatusCode != http.StatusOK {
		return errors.New("There is an error in URI of API")
	}
	return nil
}
