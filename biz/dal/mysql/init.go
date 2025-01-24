package mysql

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/soladxy/oyasumi/biz/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var _db *gorm.DB

func DefaultDB() *gorm.DB {
	return _db
}

func InitDb() {
	hlog.CtxInfof(context.Background(), "init db")
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/umami?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv(consts.MYSQL_USER),
		os.Getenv(consts.MYSQL_PASSWORD),
		os.Getenv(consts.MYSQL_HOST),
		os.Getenv(consts.MYSQL_PORT),
	)
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds),
				logger.Config{
					SlowThreshold:             200 * time.Millisecond,
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: true,
					Colorful:                  true,
				},
			),
		},
	)
	if err != nil {
		panic(err)
	}
	_db = db
}
