package internal

import (
	"time"

	"github.com/go-co-op/gocron"
)

func FunctionScheduler(f func() error) error {
	scheduler := gocron.NewScheduler(time.UTC)
	if _, err := scheduler.Every(1).Hour().Do(f); err != nil {
		return err
	}
	scheduler.StartAsync()
	return nil
}
