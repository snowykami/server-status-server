package api

import (
	"context"
	"git.liteyuki.icu/backend/golite/hertz"
	"git.liteyuki.icu/backend/golite/logger"
	"github.com/LiteyukiStudio/go-logger/log"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/hertz-contrib/cors"
	"os"
	"server-status-be/api/backend"
	"server-status-be/api/frontend"
)

var h *server.Hertz

func init() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	h = hertz.NewHertz([]config.Option{server.WithHostPorts("0.0.0.0:" + port)}, []app.HandlerFunc{})
	// cv api 状态客户端接口
	h.Use(cors.Default())

	client := h.Group("/client", backend.BeAuth)
	{
		client.GET("/ping", func(ctx context.Context, c *app.RequestContext) { c.JSON(200, "Hello, cv") })
		client.POST("/status", backend.OnPostStatus)
		client.DELETE("/host", backend.OnDeleteHost)
	}

	// api api 前端接口
	api := h.Group("/api")
	{
		api.GET("/", func(ctx context.Context, c *app.RequestContext) { c.JSON(200, "Hello, api") })
		api.GET("/status", frontend.OnGetServerStatus)
	}

	// 静态文件
	h.GET("/*file", frontend.OnGetStaticFile)
	//h.StaticFS("/", fs)
}

func Run() {

	go func() {
		err := h.Run()
		if err != nil {
			logger.Error("Run error: ", err)
		}
	}()
	log.Info("Server running...")
	select {}
}
