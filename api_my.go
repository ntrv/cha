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

func (c Client) MyStatus() (status cw.Status, err error) {
	return c.MyStatusContext(context.Background())
}
