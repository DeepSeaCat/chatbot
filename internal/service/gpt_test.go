package service

import (
	"bytes"
	"chatbot/internal/pkg/gpt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGptSendMsg(t *testing.T) {
	apikey := "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"
	reqGptBody := gpt.CreateGptReqBody("你好")
	httpClient := http.DefaultClient
	req, err := http.NewRequest("POST", BASEURL+"completions", bytes.NewBuffer([]byte(reqGptBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apikey)
	response, err := httpClient.Do(req)
	if err != nil {
		t.Errorf("error err: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		t.Errorf("gtp api status code not equals 200,code is %d", response.StatusCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("error reading body: %v", err)
	}

	gptResponseBody := gpt.ToGptResp(body)
	var reply string
	if len(gptResponseBody.Choices) > 0 {
		reply = gptResponseBody.Choices[0].Text
	}
	print(reply)
}
