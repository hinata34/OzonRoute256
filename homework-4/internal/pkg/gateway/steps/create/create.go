package create

import (
	"context"
	"homework-4/internal/model"
	"time"
)

func New() *Implementation {
	return &Implementation{}
}

type Implementation struct {
}

func (i *Implementation) Process(ctx context.Context, order *model.Order) (*model.Order, error) {

	order.States = append(order.States, model.OrderTracking{State: model.OrderStateCreate, Start: time.Now().UTC()})

	return order, nil
}

func (i *Implementation) Pipeline(ctx context.Context, stream <-chan *model.Order, workerID model.WorkerID) <-chan *model.Order {
	result := make(chan *model.Order)
	go func() {
		defer close(result)
		for order := range stream {
			createdOrder, _ := i.Process(ctx, order)
			createdOrder.WorkerID = workerID
			select {
			case <-ctx.Done():
				return
			case result <- createdOrder:
			}
		}
	}()
	return result
}
