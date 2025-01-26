// Code generated by hertz generator. DO NOT EDIT.

package oyasumi

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/soladxy/oyasumi/biz/container"
	handler "github.com/soladxy/oyasumi/biz/handler"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz, c *container.Container) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		_api.GET("/checkSsl", append(_checksslMw(), wrapHandler(handler.CheckSSL, c))...)
		_api.GET("/hello", append(_hellomethodMw(), wrapHandler(handler.HelloMethod, c))...)
	}
}

func wrapHandler(h func(context.Context, *app.RequestContext, *container.Container), container *container.Container) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		h(c, ctx, container)
	}
}
