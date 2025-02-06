package app

import (
	"context"
	"sync"

	"go-backend-template/core/kernel"
)

type Module interface {
	Info() string
	Init(engine *kernel.Engine) error
	Load(engine *kernel.Engine) error
	Start(engine *kernel.Engine) error
	Stop(wg *sync.WaitGroup, ctx context.Context) error
}
