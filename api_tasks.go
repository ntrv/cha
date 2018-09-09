package cha

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	cw "github.com/griffin-stewie/go-chatwork"
)

type Task struct {
	TaskID            int        `json:"task_id"`
	Account           cw.Account `json:"account"`
	AssignedByAccount cw.Account `json:"assigned_by_account"`
	MessageID         string     `json:"message_id"`
	Body              string     `json:"body"`
	LimitTime         time.Time  `json:"limit_time"`
	Status            string     `json:"status"`
}

func (t *Task) UnmarshalJSON(data []byte) error {
	var task cw.Task

	err = json.Unmarshal(data, &task)
	if err != nil {
		return err
	}

	t.TaskID = task.TaskID
	t.Account = task.Account
	t.AssignedByAccount = task.AssignedByAccount
	t.MessageID = task.MessageID
	t.Body = task.Body
	t.LimitTime = task.LimitDate()
	t.Status = task.Status

	return nil
}

func (c Client) RoomTasksContext(
	ctx context.Context,
	roomId string,
	params map[string]string,
) (tasks []Task, err error) {
	res, err := c.get(ctx, fmt.Sprintf("/rooms/%s/tasks", roomId), params)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &tasks)
	return
}

func (c Client) RoomTasks(roomId string, params map[string]string) ([]Task, error) {
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

func (c Client) RoomTaskContext(ctx context.Context, roomId, taskId string) (task Task, err error) {
	res, err := c.get(ctx, fmt.Sprintf("/rooms/%s/tasks/%s", roomId, taskId), map[string]string{})
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &task)
	return
}

func (c Client) RoomTask(roomId, taskId string) (Task, error) {
	return c.RoomTaskContext(context.Background(), roomId, taskId)
}
