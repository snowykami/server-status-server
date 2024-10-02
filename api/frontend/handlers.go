package frontend

import (
	"context"
	"embed"
	"github.com/LiteyukiStudio/go-logger/log"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gabriel-vasile/mimetype"
	"server-status-be/dao"
	"strings"
)

//go:embed web/*
var webFiles embed.FS

var fs2 app.FS

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
	file := c.Param("file")
	var fp string

	if file == "" {
		fp = "web/index.html"
	} else {
		fp = "web/" + file
	}

	data, err := webFiles.ReadFile(fp)
	if err != nil {
		log.Error("Read file error: ", err)
	}
	contentType := mimetype.Detect(data).String()
	if strings.HasSuffix(fp, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(fp, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(fp, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(fp, ".json") {
		contentType = "application/json"
	} else if strings.HasSuffix(fp, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(fp, ".jpg") {
		contentType = "image/jpeg"
	} else if strings.HasSuffix(fp, ".ico") {
		contentType = "image/x-icon"
	} else if strings.HasSuffix(fp, ".svg") {
		contentType = "image/svg+xml"
	}

	log.Info("Get file: ", fp, " with content type: ", contentType)
	c.Data(200, contentType, data)
}
