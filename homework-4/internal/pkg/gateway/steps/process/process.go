package process

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
	order.StorageID = model.StorageID(order.ProductID % 2)

	order.States = append(order.States, model.OrderTracking{State: model.OrderStateProcess, Start: time.Now().UTC()})

	return order, nil
}

func (i *Implementation) Pipeline(ctx context.Context, stream <-chan *model.Order) <-chan *model.Order {
	result := make(chan *model.Order)
	go func() {
		defer close(result)
		for order := range stream {
			processedOrder, _ := i.Process(ctx, order)
			select {
			case <-ctx.Done():
				return
			case result <- processedOrder:
			}
		}
	}()

	return result
}
