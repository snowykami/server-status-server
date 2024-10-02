package backend

import (
	"context"
	"encoding/json"
	"github.com/LiteyukiStudio/go-logger/log"
	"github.com/cloudwego/hertz/pkg/app"
	"server-status-be/dao"
	"server-status-be/service"
)

// be api 后端接口
func fmtReportKey(id string) string {
	return "report." + id
}

func BeAuth(ctx context.Context, c *app.RequestContext) {
	// 获取Authorization Bearer Token
	token := string(c.GetHeader("Authorization"))
	if token == "" {
		c.JSON(401, "No Token")
		c.Abort()
		return
	} else {
		if service.Token != token {
			c.JSON(401, "Invalid Token: "+token)
			c.Abort()
			return
		} else {
			c.Next(ctx)
		}
	}
}

func OnPostStatus(ctx context.Context, c *app.RequestContext) {
	report := dao.Report{}
	//err := c.BindJSON(&report)
	err := json.Unmarshal(c.GetRawData(), &report)
	log.Info(report)
	if err != nil {
		log.Error("Invalid JSON: ", err)
		c.JSON(400, "Invalid JSON: "+err.Error())
		return
	}
	dao.Save(fmtReportKey(report.Meta.ID), report)
}

func OnDeleteHost(ctx context.Context, c *app.RequestContext) {
	id, _ := c.GetPostForm("id")
	if id == "" {
		c.JSON(400, "Invalid ID")
		return
	}
	err := dao.Delete(fmtReportKey(id))
	if err != nil {
		c.JSON(400, "Delete Error: "+err.Error())
		return
	}
	c.JSON(200, "Delete Success")
}
