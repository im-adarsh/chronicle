package chronicle

import (
	"context"
	"io"
	"time"
)

type Chronicle interface {
	Register(ctx context.Context, name string, duration time.Duration, cron Cron) error
	Start() error
	io.Closer
}

type Cron func() error
