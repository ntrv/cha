package cha

import (
	"context"
	"encoding/json"

	cw "github.com/griffin-stewie/go-chatwork"
)

func (c Client) MyStatusContext(ctx context.Context) (status cw.Status, err error) {
	res, err := c.get(ctx, "/my/status", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &status)
	return
}

func (c Client) MyStatus() (cw.Status, error) {
	return c.MyStatusContext(context.Background())
}

func (c Client) MyTasksContext(
	ctx context.Context,
	params map[string]string,
) (tasks []cw.MyTask, err error) {
	res, err := c.get(ctx, "/my/tasks", params)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &tasks)
	return
}

func (c Client) MyTasks(params map[string]string) ([]cw.MyTask, error) {
	return c.MyTasksContext(context.Background(), params)
}

func (c Client) ContactsContext(ctx context.Context) (contacts []cw.Contact, err error) {
	res, err := c.get(ctx, "/contacts", map[string]string{})
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &contacts)
	return
}

func (c Client) Contacts() ([]cw.Contact, error) {
	return c.ContactsContext(context.Background())
}
