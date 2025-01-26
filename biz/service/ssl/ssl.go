package ssl

import (
	"context"
	"github.com/soladxy/oyasumi/biz/container"
)

type ISsl interface {
	HostExpired(ctx context.Context, host, port string) (bool, string, error)
}

func NewSslService(c *container.Container) ISsl {
	return newSsl(c)
}
