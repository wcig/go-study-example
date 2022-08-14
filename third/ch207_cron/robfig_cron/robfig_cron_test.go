package robfig_cron

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

type TickerJob struct{}

func (t *TickerJob) Run() {
	fmt.Println(">> run", time.Now().Format(time.StampMilli))
}

func TestFirst(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	c.Schedule(cron.Every(time.Second), &TickerJob{})
	c.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)

	// Output:
	// >> run Aug 14 22:33:39.001
	// >> run Aug 14 22:33:40.001
	// >> run Aug 14 22:33:41.001
}
