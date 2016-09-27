package gockertest

import (
	"testing"

	redis "gopkg.in/redis.v4"
)

func TestRun(t *testing.T) {
	args := Arguments{
		Ports: map[int]int{6380: 6379},
	}
	cli := Run("redis:3.2-alpine", args)
	defer cli.Cleanup()

	if cli == nil {
		t.FailNow()
	}

	ops := redis.Options{
		Addr: "localhost:6380",
	}
	rcli := redis.NewClient(&ops)

	res, err := rcli.Ping().Result()
	if err != nil {
		t.FailNow()
	}

	if res == "" {
		t.FailNow()
	}
}
