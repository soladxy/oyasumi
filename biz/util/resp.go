package util

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/soladxy/oyasumi/biz/consts"
	"net/http"
)

// SendErrResponse  pack error response
func SendErrResponse(c *app.RequestContext, respError *consts.RespError) {
	c.JSON(http.StatusBadRequest, respError)
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(c *app.RequestContext, data interface{}) {
	c.JSON(http.StatusOK, data)
}
