package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func LogMw(ctx context.Context, c *app.RequestContext) {
	uri := string(c.Request.RequestURI())
	hlog.CtxInfof(ctx, "[LogMw] Request |  Method: %s, Uri: %s, Body: %s", string(c.Request.Method()), uri, string(c.Request.Body()))
	c.Next(ctx)
	hlog.CtxInfof(ctx, "[LogMw] Response |  Uri: %s, Body: %s", uri, string(c.Response.Body()))
}
