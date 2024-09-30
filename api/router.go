package api

import (
	"context"
	"git.liteyuki.icu/backend/golite/hertz"
	"git.liteyuki.icu/backend/golite/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
)

var h *server.Hertz

func init() {
	h = hertz.NewHertz([]config.Option{server.WithHostPorts(":8088")}, []app.HandlerFunc{})
	h.GET("/", func(ctx context.Context, c *app.RequestContext) { c.JSON(200, "Hello, status be") })
	// cv api 状态客户端接口
	be := h.Group("/be", BeAuth)
	{
		be.GET("/", func(ctx context.Context, c *app.RequestContext) { c.JSON(200, "Hello, be") })
		be.POST("/report", func(ctx context.Context, c *app.RequestContext) {})
	}

	// fe api 前端接口
	fe := h.Group("/fe")
	{
		fe.GET("/", func(ctx context.Context, c *app.RequestContext) { c.JSON(200, "Hello, fe") })
	}
}

func Run() {
	go func() {
		err := h.Run()
		if err != nil {
			logger.Error("Run error: ", err)
		}
	}()

	select {}
}
