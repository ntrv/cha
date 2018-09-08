package cha

import (
	"context"
	"encoding/json"

	cw "github.com/griffin-stewie/go-chatwork"
)

func (c Client) MeContext(ctx context.Context) (me cw.Me, err error) {
	res, err := c.get(ctx, "/me", map[string]string{})
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &me)
	return
}

func (c Client) Me() (me cw.Me, err error) {
	return c.MeContext(context.Background())
}
