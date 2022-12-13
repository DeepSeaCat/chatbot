package service

import (
	"bytes"
	pb "chatbot/api/helloworld/v1"
	"chatbot/internal/conf"
	"chatbot/internal/pkg/gpt"
	"context"
	"errors"
	"fmt"
	"github.com/CatchZeng/dingtalk/pkg/dingtalk"
	"io/ioutil"
	"net/http"
	"strings"
)

const BASEURL = "https://api.openai.com/v1/"

type GreeterService struct {
	pb.UnimplementedGreeterServer
	httpClient  *http.Client
	apikey      string
	dingClient  *dingtalk.Client
	UserService UserServiceInterface
}

func NewGreeterService(gpt string, ding *conf.Bootstrap_DING, user UserServiceInterface) *GreeterService {
	dingClient := dingtalk.NewClient(ding.GetAccessToken(), ding.GetSecret())
	return &GreeterService{
		httpClient:  http.DefaultClient,
		apikey:      gpt,
		dingClient:  dingClient,
		UserService: user,
	}
}

func (s *GreeterService) ChatMsg(ctx context.Context, req *pb.DingRequest) (*pb.DingReply, error) {
	reqGptBody := gpt.CreateGptReqBody(req.Text.Content)
	reply, err := s.GptSendMsg(ctx, reqGptBody)
	if err != nil {
		s.DingSendMsg("机器人太累了，让她休息会儿，过一会儿再来请求。", []string{req.SenderNick})
		return nil, nil
	}
	if reply == "" {
		return nil, nil
	}
	if s.UserService.ClearUserSessionContext(req.SenderID, req.Text.Content) {
		s.DingSendMsg("上下文已经清空了，你可以问下一个问题啦。", []string{req.SenderNick})
		return nil, nil
	}
	s.UserService.SetUserSessionContext(req.SenderID, req.Text.Content, reply)
	s.DingSendMsg(reply, []string{req.SenderNick})
	return &pb.DingReply{}, nil
}

func (s *GreeterService) GptSendMsg(ctx context.Context, requestData string) (string, error) {
	req, err := http.NewRequest("POST", BASEURL+"completions", bytes.NewBuffer([]byte(requestData)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apikey)
	response, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("gtp api status code not equals 200,code is %d", response.StatusCode))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := gpt.ToGptResp(body)
	var reply string
	if len(gptResponseBody.Choices) > 0 {
		reply = gptResponseBody.Choices[0].Text
	}
	return reply, nil

}

func (s GreeterService) DingSendMsg(msg string, at []string) error {
	msg = strings.TrimPrefix(msg, ".")
	msg = strings.TrimPrefix(msg, "\n")
	msg = strings.TrimPrefix(msg, "\n")
	dingMsg := dingtalk.NewTextMessage().SetContent(msg).SetAt(at, false)
	_, _, err := s.dingClient.Send(dingMsg)
	if err != nil {
		return err
	}
	return nil
}
