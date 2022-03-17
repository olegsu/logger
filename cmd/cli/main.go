package main

import (
	"time"

	"github.com/olegsu/logger"
)

func main() {
	lgr := logger.New()
	lgr.Info("hey")
	time.Sleep(10 * time.Second)
}
