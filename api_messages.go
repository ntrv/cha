package cha

import (
	"context"
	"encoding/json"
	"fmt"

	cw "github.com/griffin-stewie/go-chatwork"
)

func (c Client) RoomMessagesContext(
	ctx context.Context,
	roomId string,
	params map[string]string,
) (msgs []cw.Message, err error) {
	res, err := c.get(ctx, fmt.Sprintf("/rooms/%s/messages", roomId), params)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &msgs)
	return
}

func (c Client) RoomMessages(roomId string, params map[string]string) ([]cw.Message, error) {
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
) (msg cw.Message, err error) {
	res, err := c.get(ctx, fmt.Sprintf("/rooms/%s/messages/%s", roomId, messageId), map[string]string{})
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &msg)
	return
}

func (c Client) RoomMessage(roomId, messageId string) (cw.Message, error) {
	return c.RoomMessageContext(context.Background(), roomId, messageId)
}
