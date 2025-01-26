package mysql

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/soladxy/oyasumi/biz/config"
	"github.com/soladxy/oyasumi/biz/dal/mysql/model"
	"github.com/soladxy/oyasumi/biz/dal/mysql/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type _db struct {
	native *gorm.DB     // 原生gorm
	gen    *query.Query // gorm_gen
}

// NewGenDBClient gorm_gen 使用，请勿在其他场景使用
func NewGenDBClient(c *config.Config) (*gorm.DB, error) {
	db, err := newMySQL(c)
	if err != nil {
		return nil, err
	}
	return db.native, nil
}

func newMySQL(c *config.Config) (*_db, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/umami?charset=utf8mb4&parseTime=True&loc=Local",
		c.MySQL.User, c.MySQL.Password, c.MySQL.Host, c.MySQL.Port,
	)
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: newGormLogger(),
		},
	)
	if err != nil {
		return nil, err
	}
	return &_db{
		native: db,
		gen:    query.Use(db),
	}, nil
}

var _ IMySQL = (*_db)(nil)

func (db *_db) GetLatestWebsiteRecord(ctx context.Context) (*model.Website, error) {
	w := db.gen.Website
	last, err := w.WithContext(ctx).Last()
	if err != nil {
		hlog.CtxErrorf(ctx, "[GetLatestRecord] get last record err: %v", err)
		return nil, err
	}
	return last, nil
}
