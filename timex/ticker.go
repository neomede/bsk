package timex

import (
	"context"
	"time"

	"github.com/socialpoint-labs/bsk/server"
)

// RunInterval runs the provided function at intervals specified by the interval argument.
func RunInterval(ctx context.Context, interval time.Duration, f func()) {
	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for {
		f()

		select {
		case <-ticker.C:
			// continue
		case <-ctx.Done():
			return
		}
	}
}

// IntervalRunner returns a server.Runner that runs the function RunInterval with the provided context
func IntervalRunner(interval time.Duration, f func()) server.Runner {
	return server.RunnerFunc(func(ctx context.Context) {
		RunInterval(ctx, interval, f)
	})
}
