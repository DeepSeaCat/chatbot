package gpt

import "encoding/json"

// ChatGPTReq 请求体
type ChatGPTReq struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
}

func CreateGptReqBody(msg string) string {
	req := &ChatGPTReq{
		Model:            "text-davinci-003",
		Prompt:           msg,
		MaxTokens:        1024,
		Temperature:      0.7,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	requestData, _ := json.Marshal(req)
	return string(requestData)
}
