package model

import "time"

// 消息对象
type MessageObj struct {
	// 消息ID
	ID string

	// 消息信息
	MsgBody interface{}

	// 当前时间
	NowTime time.Time
}

// 创建新的消息对象
// id：消息Id
// msgObj:消息体
// 返回值：
// 消息对象
func NewMessageObj(id string, msgBody interface{}) *MessageObj {
	return &MessageObj{
		ID:      id,
		MsgBody: msgBody,
		NowTime: time.Now(),
	}
}
