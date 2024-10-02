package frontend

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"os"
	"server-status-be/dao"
	"strings"
)

func OnGetServerStatus(ctx context.Context, c *app.RequestContext) {
	ret := make(map[string]interface{})
	for k, v := range dao.GetAll() {
		if strings.HasPrefix(k, "report.") {
			reportName := strings.TrimPrefix(k, "report.")
			ret[reportName] = v
		}
	}
	c.JSON(200, ret)
	ctx.Done()
}

func OnGetStaticFile(ctx context.Context, c *app.RequestContext) {
	c.Request.URI().Path()
	file := c.Param("file")
	fp := "./web/" + file

	if file == "" {
		fp = "web/index.html"
	}
	// 判断文件是否存在
	if _, err := os.Stat(fp); err != nil {
		fp += ".html"
		if _, err := os.Stat(fp); err != nil {
			c.JSON(404, "File Not Found")
			return
		} else {
			c.File(fp)
			return
		}
	} else {
		c.File(fp)
	}
	ctx.Done()
}
