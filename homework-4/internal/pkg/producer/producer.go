package producer

import (
	"context"
	"homework-4/configs"
	"homework-4/internal/model"
)

func Orders(ctx context.Context) <-chan *model.Order {
	result := make(chan *model.Order)

	go func() {
		defer close(result)

		for i := model.ProductID(1); i <= configs.OrdersTotal; i++ {
			p := new(model.Order)
			p.ProductID = i
			select {
			case <-ctx.Done():
				return
			case result <- p:
			}
		}
	}()

	return result
}
