package main

import (
	"fmt"

	"github.com/takecy/gockertest"
	redis "gopkg.in/redis.v5"
)

// Example: run redis container from private registry
// $ go run yours.go
func main() {
	fmt.Printf("start.example\n")

	args := gockertest.Arguments{
		Ports:        map[int]int{6379: 6379},
		RequireLogin: true, // require basic authentication
		Login: gockertest.Login{
			User:     "yourname",          // change to your username
			Password: "pass",              // change to your password
			Registry: "registry.yours.io", // change to your registy domain
		},
	}
	cli := gockertest.Run("registry.yours.io/redis:3.2-alpine", args)
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
