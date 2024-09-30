package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"server-status-be/dao"
)

// be api 后端接口

func BeAuth(ctx context.Context, c *app.RequestContext) {
	// 获取Authorization Bearer Token
	token := string(c.GetHeader("Authorization"))
	if token == "" {
		c.JSON(401, "Unauthorized")
		c.Abort()
		return
	}
}

func OnPostReport(ctx context.Context, c *app.RequestContext) {
	report := dao.Report{}
	err := c.BindJSON(&report)
	if err != nil {
		c.JSON(400, "Invalid JSON")
		return
	}
}
