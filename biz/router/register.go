// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/soladxy/oyasumi/biz/container"
	oyasumi "github.com/soladxy/oyasumi/biz/router/oyasumi"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz, c *container.Container) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	oyasumi.Register(r, c)
}
