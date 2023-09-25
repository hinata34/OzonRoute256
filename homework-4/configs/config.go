package configs

import "homework-4/internal/model"

type FanOutType uint64

const (
	OrdersTotal  = model.ProductID(20)
	WorkersTotal = model.WorkerID(4)
	FanOutTotal  = FanOutType(5)
)
