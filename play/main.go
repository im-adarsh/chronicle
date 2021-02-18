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
	c.Register(ctx, "say hello", time.Second*2, printHello)
	c.Register(ctx, "say world", time.Second*4, printWorld)
	c.Start()
}

func printHello() error {
	fmt.Println("hello")
	return nil
}

func printWorld() error {
	fmt.Println("world")
	return nil
}
