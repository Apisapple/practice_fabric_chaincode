package entity

import "encoding/json"

type Message struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

func (m *Message) MakeMessage(title, content string) *Message {
	return &Message{Title: title, Content: content}
}

func (m *Message) ToObject(data []byte) error {

	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}

	return nil
}

func (m *Message) ToJson() ([]byte, error) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
