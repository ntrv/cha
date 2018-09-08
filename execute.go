package cha

import (
	"context"
	"net/http"

	"golang.org/x/net/context/ctxhttp"
)

func (c Client) execute(
	ctx context.Context,
	method, spath string,
	params map[string]string,
) ([]byte, error) {
	var res *http.Response

	req, err := c.newRequest(method, spath, params)
	if err != nil {
		return []byte(``), err
	}

	res, err = ctxhttp.Do(ctx, c.HTTPClient, req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return []byte(``), err
	}

	return utils.writeRes(res)
}

func (c Client) Get(
	ctx context.Context,
	spath string,
	params map[string]string,
) ([]byte, error) {
	return c.execute(ctx, http.MethodGet, spath, params)
}

func (c Client) Post(
	ctx context.Context,
	spath string,
	params map[string]string,
) ([]byte, error) {
	return c.execute(ctx, http.MethodPost, spath, params)
}

func (c Client) Put(
	ctx context.Context,
	spath string,
	params map[string]string,
) ([]byte, error) {
	return c.execute(ctx, http.MethodPut, spath, params)
}

func (c Client) Delete(
	ctx context.Context,
	spath string,
	params map[string]string,
) ([]byte, error) {
	return c.execute(ctx, http.MethodDelete, spath, params)
}
