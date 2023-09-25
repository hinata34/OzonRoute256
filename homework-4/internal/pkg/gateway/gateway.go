package gateway

import (
	"context"
	"homework-4/configs"
	"homework-4/internal/model"
	completestep "homework-4/internal/pkg/gateway/steps/complete"
	createstep "homework-4/internal/pkg/gateway/steps/create"
	processstep "homework-4/internal/pkg/gateway/steps/process"
	"sync"
)

func New(create *createstep.Implementation, process *processstep.Implementation, complete *completestep.Implementation) *Implementation {
	return &Implementation{
		create:   create,
		process:  process,
		complete: complete,
	}
}

type Implementation struct {
	create   *createstep.Implementation
	process  *processstep.Implementation
	complete *completestep.Implementation
}

func (i *Implementation) Pipeline(ctx context.Context, inStream <-chan *model.Order, outStream chan<- *model.Order, workerID model.WorkerID) {
	createdCh := i.create.Pipeline(ctx, inStream, workerID)

	fanOutProcess := make([]<-chan *model.Order, configs.FanOutTotal)
	for it := configs.FanOutType(0); it < configs.FanOutTotal; it++ {
		fanOutProcess[it] = i.process.Pipeline(ctx, createdCh)
	}

	completed := i.complete.Pipeline(ctx, fanIn(ctx, fanOutProcess))
	for completedOrder := range completed {
		select {
		case <-ctx.Done():
			return
		case outStream <- completedOrder:
		}
	}
}

func fanIn(ctx context.Context, chans []<-chan *model.Order) <-chan *model.Order {
	muliteplexed := make(chan *model.Order)

	var wg sync.WaitGroup
	for _, ch := range chans {
		wg.Add(1)

		go func(ch <-chan *model.Order) {
			defer wg.Done()
			for v := range ch {
				select {
				case <-ctx.Done():
					return
				case muliteplexed <- v:
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(muliteplexed)
	}()

	return muliteplexed
}
