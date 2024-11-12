package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Djone4213/telegramSender/entity"
	"net/http"
)

const (
	URL = "https://api.telegram.org/bot"
)

//type Sender interface {
//	SendMessage(msg entity.MessageText) (entity.OutputBody, error)
//	SendVideo(msg entity.VideoMessage) (entity.OutputBody, error)
//	EditMessageText(msg entity.EditMessageText) (entity.OutputBody, error)
//	DeleteMessage(msg entity.DeleteMessage) (entity.OutputBody, error)
//	EditMessageReplyMarkup(msg entity.EditMessageReplyMarkup) (entity.OutputBody, error)
//}

type sender struct {
	token string
}

func NewSender(token string) *sender {
	return &sender{
		token: token,
	}
}

func (s *sender) SendMessage(msg entity.MessageText) (entity.OutputBodyMessage, error) {
	var outMsg entity.OutputBodyMessage

	if msg.ChatId == 0 {
		return outMsg, errors.New("Not chatId")
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return outMsg, err
	}

	result, err := s.post(jsonMsg, fmt.Sprintf("%s%s/sendMessage", URL, s.token))
	if err != nil {
		return outMsg, err
	}

	err = json.Unmarshal(result, &outMsg)
	if err != nil {
		return outMsg, err
	}

	return outMsg, nil
}

func (s *sender) SendVideo(msg entity.VideoMessage) (entity.OutputBodyMessage, error) {
	var outMsg entity.OutputBodyMessage

	if msg.ChatId == 0 {
		return outMsg, errors.New("Not chatId")
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return outMsg, err
	}

	result, err := s.post(jsonMsg, fmt.Sprintf("%s%s/sendVideo", URL, s.token))
	if err != nil {
		return outMsg, err
	}

	err = json.Unmarshal(result, &outMsg)
	if err != nil {
		return outMsg, err
	}

	return outMsg, nil
}

func (s *sender) EditMessageText(msg entity.EditMessageText) (entity.OutputBodyMessage, error) {
	var outMsg entity.OutputBodyMessage

	if msg.ChatId == 0 {
		return outMsg, errors.New("Not chatId")
	}

	if msg.MessageId == 0 {
		return outMsg, errors.New("Not messageId to replace")
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return outMsg, err
	}

	result, err := s.post(jsonMsg, fmt.Sprintf("%s%s/editMessageText", URL, s.token))
	if err != nil {
		return outMsg, err
	}

	err = json.Unmarshal(result, &outMsg)
	if err != nil {
		return outMsg, err
	}

	return outMsg, nil
}

func (s *sender) DeleteMessage(msg entity.DeleteMessage) (entity.OutputBodyDeleteMessage, error) {
	var outMsg entity.OutputBodyDeleteMessage

	if msg.ChatId == 0 {
		return outMsg, errors.New("Not chatId")
	}

	if msg.MessageId == 0 {
		return outMsg, errors.New("Not messageId to delete")
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return outMsg, err
	}

	result, err := s.post(jsonMsg, fmt.Sprintf("%s%s/deleteMessage", URL, s.token))
	if err != nil {
		return outMsg, err
	}

	err = json.Unmarshal(result, &outMsg)
	if err != nil {
		return outMsg, err
	}

	return outMsg, nil
}

func (s *sender) EditMessageReplyMarkup(msg entity.EditMessageReplyMarkup) (entity.OutputBodyMessage, error) {
	var outMsg entity.OutputBodyMessage

	if *msg.ChatId == 0 {
		return outMsg, errors.New("Not chatId")
	}

	if *msg.MessageId == 0 {
		return outMsg, errors.New("Not messageId to delete")
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return outMsg, err
	}

	result, err := s.post(jsonMsg, fmt.Sprintf("%s%s/editMessageReplyMarkup", URL, s.token))
	if err != nil {
		return outMsg, err
	}

	err = json.Unmarshal(result, &outMsg)
	if err != nil {
		return outMsg, err
	}

	return outMsg, nil
}

func (s *sender) post(msg []byte, url string) ([]byte, error) {

	var result []byte
	var outMsg map[string]interface{}
	var err error

	jsonStr := bytes.NewReader(msg)

	req, err := http.NewRequest("POST", url, jsonStr)
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

	if resp.StatusCode != http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&outMsg)
		if err != nil {
			return nil, err
		}

		result, err = json.Marshal(outMsg)
		if err != nil {
			return nil, err
		}

		return result, errors.New("Error from Telegram server")
	}

	err = json.NewDecoder(resp.Body).Decode(&outMsg)
	if err != nil {
		return nil, err
	}

	result, err = json.Marshal(outMsg)
	if err != nil {
		return nil, err
	}

	return result, nil
}
