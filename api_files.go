package cha

import (
	"context"
	"encoding/json"
	"fmt"

	cw "github.com/griffin-stewie/go-chatwork"
)

func (c Client) RoomFilesContext(
	ctx context.Context,
	roomId string,
	params map[string]string,
) (files []cw.File, err error) {
	res, err := c.get(ctx, fmt.Sprintf("/rooms/%s/files", roomId), params)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &files)
	return
}

func (c Client) RoomFiles(roomId string, params map[string]string) ([]cw.File, error) {
	return c.RoomFilesContext(context.Background(), roomId, params)
}

func (c Client) RoomFileContext(
	ctx context.Context,
	roomId, fileId string,
	params map[string]string,
) (file cw.File, err error) {
	res, err := c.get(ctx, fmt.Sprintf("/rooms/%s/files/%s", roomId, fileId), params)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &file)
	return
}
