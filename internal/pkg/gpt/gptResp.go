package gpt

import "encoding/json"

// ChatGPTResp 响应体
type ChatGPTResp struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChoiceItem           `json:"choices"`
	Usage   map[string]interface{} `json:"usage"`
}

func ToGptResp(msg []byte) *ChatGPTResp {
	resp := &ChatGPTResp{}
	json.Unmarshal(msg, resp)
	return resp
}
