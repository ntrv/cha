package cha

import (
	"context"
	"encoding/json"

	cw "github.com/griffin-stewie/go-chatwork"
)

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
