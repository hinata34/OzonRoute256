package complete

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
	order.PickUpPointID = model.PickUpPointID(order.ProductID + model.ProductID(order.StorageID))

	order.States = append(order.States, model.OrderTracking{State: model.OrderStateComplete, Start: time.Now().UTC()})

	return order, nil
}

func (i *Implementation) Pipeline(ctx context.Context, stream <-chan *model.Order) <-chan *model.Order {
	result := make(chan *model.Order)
	go func() {
		defer close(result)
		for order := range stream {
			completedOrder, _ := i.Process(ctx, order)
			select {
			case <-ctx.Done():
				return
			case result <- completedOrder:
			}
		}
	}()

	return result
}
