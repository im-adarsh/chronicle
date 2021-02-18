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
	c.Register(ctx, "worker_1", time.Second*2, doTask1)
	c.Register(ctx, "worker_2", time.Second*4, doTask2)
	c.Register(ctx, "worker_3", time.Second*6, doTask3)

	c.Start()
	defer c.Close()
}

func doTask1() error {
	fmt.Println("doing task 1")
	return nil
}

func doTask2() error {
	fmt.Println("doing task 2")
	return nil
}

func doTask3() error {
	fmt.Println("doing task 3")
	return nil
}
