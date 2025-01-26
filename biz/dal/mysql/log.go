package mysql

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm/logger"
	"time"
)

var _ logger.Interface = (*gormLogger)(nil)

type gormLogger struct {
}

func newGormLogger() logger.Interface {
	return &gormLogger{}
}

func (g *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g *gormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	hlog.CtxInfof(ctx, s, i...)
}

func (g *gormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	hlog.CtxWarnf(ctx, s, i...)
}

func (g *gormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	hlog.CtxErrorf(ctx, s, i...)
}

func (g *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	hlog.CtxTracef(ctx, "sql: %s, rowsAffected: %d, begin: %s, err: %v", sql, rowsAffected, begin.Format(time.DateTime), err)
}
