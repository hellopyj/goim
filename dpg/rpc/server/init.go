package server

import (
	"github.com/Terry-Mao/goim/dpg/rpc/server/handler"
	"github.com/Terry-Mao/goim/dpg/rpc/server/push"
	"github.com/Terry-Mao/goim/internal/logic"
	"github.com/Terry-Mao/goim/internal/logic/conf"
)

func New(c *conf.Config,l *logic.Logic)  {
	push.New(c,l)
	handler.New(c)
}

