package main

import (
	"fmt"

	"github.com/takecy/gockertest"
	redis "gopkg.in/redis.v4"
)

// Example: run redis container
// $ go run redis.go
// after complete, run
// $ docker ps -a
// not to leave garbage.
func main() {
	fmt.Printf("start.example\n")

	args := gockertest.Arguments{
		Ports: map[int]int{6379: 6379},
	}
	cli := gockertest.Run("redis:3.2-alpine", args)
	defer cli.Cleanup()

	fmt.Printf("started.container: %s\n", cli.ID)

	ops := redis.Options{
		Addr: "localhost:6379",
	}
	rcli := redis.NewClient(&ops)

	fmt.Printf("init.redis.client: %s\n", ops.Addr)

	res, err := rcli.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("redis.ping: %s\n", res)
}
