package kernel

import (
	"context"

	"github.com/flamego/flamego"
	"go-backend-template/core/store/mysql"
)

type Engine struct {
	Ctx    context.Context
	Cancel context.CancelFunc
	Fg     *flamego.Flame
	Mysql  *mysql.Orm
}

var Kernel *Engine
