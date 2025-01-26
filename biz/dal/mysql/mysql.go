package mysql

import (
	"context"
	"github.com/soladxy/oyasumi/biz/config"
	"github.com/soladxy/oyasumi/biz/dal/mysql/model"
)

type IMySQL interface {
	GetLatestWebsiteRecord(ctx context.Context) (*model.Website, error)
}

func NewMySQL(c *config.Config) (IMySQL, error) {
	return newMySQL(c)
}
