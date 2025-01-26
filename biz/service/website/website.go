package website

import (
	"context"
	"github.com/soladxy/oyasumi/biz/container"
)

type IWebsite interface {
	GetLatestWebsiteInfo(ctx context.Context) (string, error)
}

func NewWebsiteService(c *container.Container) IWebsite {
	return newWebsite(c)
}
