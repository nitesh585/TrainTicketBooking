package main

import (
	"context"
	"loggingService/gokafka"
)

func main() {
	for {
		ctx := context.Background()
		go gokafka.LogConsumer(ctx, "error")
		go gokafka.LogConsumer(ctx, "info")
		go gokafka.LogConsumer(ctx, "debug")
		go gokafka.LogConsumer(ctx, "warn")
	}
}
