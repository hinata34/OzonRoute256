package model

import "time"

type OrderState string

const (
	OrderStateCreate   = OrderState("create")
	OrderStateProcess  = OrderState("process")
	OrderStateComplete = OrderState("complete")
)

type ProductID uint64
type StorageID uint64
type PickUpPointID uint64
type WorkerID uint64

type Order struct {
	ProductID     ProductID
	StorageID     StorageID
	PickUpPointID PickUpPointID
	WorkerID      WorkerID
	States        []OrderTracking
}

type OrderTracking struct {
	State OrderState
	Start time.Time
}
