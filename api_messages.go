package cha

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	cw "github.com/griffin-stewie/go-chatwork"
)

type Message struct {
	MessageID  string     `json:"message_id"`
	Account    cw.Account `json:"account"`
	Body       string     `json:"body"`
	SendTime   time.Time  `json:"send_time"`
	UpdateTime time.Time  `json:"update_time"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var msg cw.Message

	err := json.Unmarshal(data, &msg)
	if err != nil {
		return err
	}

	m.MessageID = msg.MessageID
	m.Account = msg.Account
	m.Body = msg.Body
	m.SendTime = msg.SendDate()
	m.UpdateTime = msg.UpdateDate()

	return nil
}

func (c Client) RoomMessagesContext(
	ctx context.Context,
	roomId string,
	params map[string]string,
) (msgs []Message, err error) {
	res, err := c.get(ctx, fmt.Sprintf("/rooms/%s/messages", roomId), params)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &msgs)
	return
}

func (c Client) RoomMessages(roomId string, params map[string]string) ([]Message, error) {
	return c.RoomMessagesContext(context.Background(), roomId, params)
}

func (c Client) PostRoomMessageContext(
	ctx context.Context,
	roomId string,
	body string,
) ([]byte, error) {
	return c.post(ctx, fmt.Sprintf("/rooms/%s/messages", roomId), map[string]string{"body": body})
}

func (c Client) PostRoomMessage(roomId string, body string) ([]byte, error) {
	return c.PostRoomMessageContext(context.Background(), roomId, body)
}

func (c Client) RoomMessageContext(
	ctx context.Context,
	roomId, messageId string,
) (msg Message, err error) {
	res, err := c.get(ctx, fmt.Sprintf("/rooms/%s/messages/%s", roomId, messageId), map[string]string{})
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &msg)
	return
}

func (c Client) RoomMessage(roomId, messageId string) (Message, error) {
	return c.RoomMessageContext(context.Background(), roomId, messageId)
}
