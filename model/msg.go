/*
# @Author ww
# @Time 2019/11/5 13:45
# @File msg.go
*/
package model

const (
	// 消息发送方式
	EMAIL = iota
	SMS
	PHONE
	// 消息类型：警告，恢复
	ALERT = iota
	RESUME
)

type Msg struct {
	Id           string `json:"id"`         // 消息id
	Ts           int64  `json:"ts"`         // 时间戳
	Type         int    `json:"type"`       // 消息类型
	Content      string `json:"content"`    // 消息内容
	Confirm      bool   `json:"confirm"`    // 消息是否确认
	Ts_confirm   int64  `json:"ts_confirm"` // 消息确认时间
	Type_confirm int    // 消息确认方式
}

func NewMsg() *Msg {
	return nil
}
