package chronicle

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type chronicle struct {
	wg    *sync.WaitGroup
	tasks map[string]task
}

type task struct {
	duration time.Duration
	cron     Cron
}

func NewChronicle() Chronicle {
	c := chronicle{}
	c.wg = new(sync.WaitGroup)
	c.tasks = map[string]task{}
	return &c
}

func (c *chronicle) Register(ctx context.Context, name string, duration time.Duration, cron Cron) error {
	_, ok := c.tasks[name]
	if ok {
		return errors.New("task is already registered")
	}
	c.tasks[name] = task{duration: duration, cron: cron}
	return nil
}

func (c *chronicle) Start() error {
	fmt.Println(fmt.Sprintf("Starting %d tasks", len(c.tasks)))
	if len(c.tasks) == 0 {
		return errors.New("no task is registered")
	}
	c.wg.Add(len(c.tasks))
	for n, v := range c.tasks {
		tt := time.NewTicker(v.duration)
		tc := make(chan bool)
		go func() {
			for {
				select {
				case t := <-tt.C:
					fmt.Println(fmt.Sprintf("[Tick %s] at %+v ", n, t))
					_ = v.cron()
				case <-tc:
					return
				}
			}
		}()
	}
	c.wg.Wait()
	return nil
}

func (c *chronicle) Close() error {
	for _, _ = range c.tasks {
		c.wg.Done()
	}
	return nil
}
