package chronicle

import (
	"context"
	"errors"
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
	c.wg = &sync.WaitGroup{}
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
	if len(c.tasks) == 0 {
		return errors.New("no task is registered")
	}
	for _, v := range c.tasks {
		err := v.cron()
		if err != nil {
			return err
		}
		c.wg.Add(1)
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
