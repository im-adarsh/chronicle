package main

import (
	"context"
	"fmt"
	"time"

	"github.com/im-adarsh/chronicle"
)

func main() {

	ctx := context.Background()
	c := chronicle.NewChronicle()
	c.Register(ctx, "", time.Minute, print)
	c.Start()
}

func print() error {
	fmt.Println("hello")
	return nil
}
