package utils

import (
	"strconv"
	"time"
)

type RateLimit struct {
	Limit     int
	Remaining int
	ResetTime time.Time
}

type HTTPHeader interface {
	Get(string) string
}

func rateLimit(header HTTPHeader) (*RateLimit, error) {
	limit, err := strconv.Atoi(header.Get("X-RateLimit-Limit"))
	if err != nil {
		return nil, err
	}

	remaining, err := strconv.Atoi(header.Get("X-RateLimit-Remaining"))
	if err != nil {
		return nil, err
	}

	resetTime, err := strconv.ParseInt(header.Get("X-RateLimit-Reset"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &RateLimit{
		Limit:     limit,
		Remaining: remaining,
		ResetTime: time.Unix(resetTime, 0),
	}, nil
}
