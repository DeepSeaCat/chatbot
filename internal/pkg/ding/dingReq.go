package ding

// SendMsg 发送的消息体
type SendMsg struct {
	Text    Text   `json:"text"`
	Msgtype string `json:"msgtype"`
}
