package cha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	cw "github.com/griffin-stewie/go-chatwork"
)

type File struct {
	FileID      int        `json:"file_id"`
	Account     cw.Account `json:"account"`
	MessageID   string     `json:"message_id"`
	Filename    string     `json:"filename"`
	Filesize    int        `json:"filesize"`
	UploadTime  time.Time  `json:"upload_time"`
	DownloadURL *url.URL   `json:"download_url"`
}

func (f *File) UnmarshalJSON(data []byte) error {
	var file cw.File

	err := json.Unmarshal(data, &file)
	if err != nil {
		return err
	}

	f.FileID = file.FileID
	f.Account = file.Account
	f.MessageID = file.MessageID
	f.Filename = file.Filename
	f.Filesize = file.Filesize
	f.UploadTime = file.UploadDate()

	f.DownloadURL, err = url.Parse(file.DownloadURL)
	if err != nil {
		return err
	}

	return nil
}

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
