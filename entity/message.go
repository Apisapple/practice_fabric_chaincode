package entity

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

func (m *Message) MakeMessage(title, content string) *Message {
	return &Message{Title: title, Content: content}
}

func (m *Message) ToObject(data []byte) error {
	fmt.Printf("input data is %s\n", data)
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}

	fmt.Printf("Convert data is %s \n", m)
	return nil
}

func (m *Message) ToJson() ([]byte, error) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

type Request struct {
	Msg *Message `json:"msg,omitempty"`
}

func (m *Request) ToObject(data []byte) error {
	fmt.Printf("input data is %s\n", data)
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}

	fmt.Printf("Convert data is %s \n", m)
	return nil
}

func (m *Request) ToJson() ([]byte, error) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
