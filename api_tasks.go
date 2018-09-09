package cha

import (
	"context"
	"encoding/json"
	"fmt"

	cw "github.com/griffin-stewie/go-chatwork"
)

func (c Client) RoomTasksContext(
	ctx context.Context,
	roomId string,
	params map[string]string,
) (tasks []cw.Task, err error) {
	res, err := c.get(ctx, fmt.Sprintf("/rooms/%s/tasks", roomId), params)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &tasks)
	return
}

func (c Client) RoomTasks(roomId string, params map[string]string) ([]cw.Task, error) {
	return c.RoomTasksContext(context.Background(), roomId, params)
}

func (c Client) PostRoomTaskContext(
	ctx context.Context,
	roomId string,
	params map[string]string,
) ([]byte, error) {
	return c.post(ctx, fmt.Sprintf("/rooms/%s/tasks", roomId), params)
}

func (c Client) PostRoomTask(roomId string, params map[string]string) ([]byte, error) {
	return c.PostRoomTaskContext(context.Background(), roomId, params)
}

func (c Client) RoomTaskContext(ctx context.Context, roomId, taskId string) (task cw.Task, err error) {
	res, err := c.get(ctx, fmt.Sprintf("/rooms/%s/tasks/%s", roomId, taskId), map[string]string{})
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &task)
	return
}

func (c Client) RoomTask(roomId, taskId string) (cw.Task, error) {
	return c.RoomTaskContext(context.Background(), roomId, taskId)
}
