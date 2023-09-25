package main

import (
	"context"
	"encoding/json"
	"homework-4/configs"
	"homework-4/internal/model"
	"homework-4/internal/pkg/gateway"
	"homework-4/internal/pkg/gateway/steps/complete"
	"homework-4/internal/pkg/gateway/steps/create"
	"homework-4/internal/pkg/gateway/steps/process"
	"homework-4/internal/pkg/producer"
	"os"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	create := create.New()
	process := process.New()
	complete := complete.New()

	handler := gateway.New(create, process, complete)

	orders := producer.Orders(ctx)
	completedOrders := make(chan *model.Order)
	var wg sync.WaitGroup

	for i := model.WorkerID(1); i <= configs.WorkersTotal; i++ {
		wg.Add(1)
		go worker(ctx, &wg, handler, orders, completedOrders, i)
	}

	go func() {
		wg.Wait()
		close(completedOrders)
	}()

	for order := range completedOrders {
		err := json.NewEncoder(os.Stdout).Encode(order)
		if err != nil {
			return
		}
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, handler *gateway.Implementation, orders <-chan *model.Order, completedOrders chan<- *model.Order, workerID model.WorkerID) {
	defer wg.Done()
	handler.Pipeline(ctx, orders, completedOrders, workerID)
}
