package cha

import (
	"context"
	"encoding/json"
	"fmt"

	cw "github.com/griffin-stewie/go-chatwork"
)

func (c Client) RoomsContext(ctx context.Context) (rooms []cw.Room, err error) {
	res, err := c.get(ctx, "/rooms", map[string]string{})
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &rooms)
	return
}

func (c Client) Rooms() ([]cw.Room, error) {
	return c.RoomsContext(context.Background())
}

func (c Client) CreateRoomContext(ctx context.Context, params map[string]string) ([]byte, error) {
	return c.post(ctx, "/rooms", params)
}

func (c Client) CreateRoom(params map[string]string) ([]byte, error) {
	return c.CreateRoomContext(context.Background(), params)
}

func (c Client) UpdateRoomContext(
	ctx context.Context,
	roomId string,
	params map[string]string,
) ([]byte, error) {
	return c.put(ctx, fmt.Sprintf("/rooms/%s", roomId), params)
}

func (c Client) UpdateRoom(roomId string, params map[string]string) ([]byte, error) {
	return c.UpdateRoomContext(context.Background(), roomId, params)
}

func (c Client) DeleteRoomContext(
	ctx context.Context,
	roomId string,
	params map[string]string,
) ([]byte, error) {
	return c.delete(ctx, fmt.Sprintf("/rooms/%s", roomId), params)
}

func (c Client) DeleteRoom(roomId string, params map[string]string) {
	return c.DeleteRoomContext(context.Background(), roomId, params)
}

func (c Client) RoomMembersContext(
	ctx context.Context,
	roomId string,
) (members []cw.Member, err error) {
	res, err := c.get(ctx, fmt.Sprintf("/rooms/%s/members", roomId), map[string]string)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &members)
	return
}

func (c Client) RoomMembers(roomId string) ([]cw.Member, error) {
	return c.RoomMembersContext(context.Background(), roomId)
}

func (c Client) UpdateRoomMembersContext(
	ctx context.Context,
	roomId string,
	params map[string]string,
) ([]byte, error) {
	return c.put(ctx, fmt.Sprintf("/rooms/%s/members"), params)
}

func (c Client) UpdateRoomMembers(roomId string, params map[string]string) ([]byte, error) {
	return c.UpdateRoomMembersContext(context.Background(), roomId, params)
}
