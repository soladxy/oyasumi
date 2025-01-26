package website

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/soladxy/oyasumi/biz/container"
	"time"
)

type _website struct {
	c *container.Container
}

func newWebsite(c *container.Container) *_website {
	return &_website{
		c: c,
	}
}

var _ IWebsite = (*_website)(nil)

func (w *_website) GetLatestWebsiteInfo(ctx context.Context) (string, error) {
	record, err := w.c.MySQL.GetLatestWebsiteRecord(ctx)
	if err != nil {
		hlog.CtxErrorf(ctx, "[GetLatestWebsiteInfo] error: %v", err)
		return "", err
	}
	return fmt.Sprintf("ID: %s, Domain: %s, CreatedAt: %s, CreatedBy: %s", record.WebsiteID, record.Domain, record.CreatedAt.Format(time.DateTime), record.CreatedBy), nil
}
