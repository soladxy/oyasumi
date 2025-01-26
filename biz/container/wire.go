//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/soladxy/oyasumi/biz/config"
	"github.com/soladxy/oyasumi/biz/dal/mysql"
	"gorm.io/gorm"
)

type Container struct {
	MySQL mysql.IMySQL
}

func InitService(path string) (*Container, error) {
	wire.Build(
		config.NewServiceConfig,
		mysql.NewMySQL,
		wire.Struct(new(Container), "*"),
	)
	return &Container{}, nil
}

func InitGormGen(path string) (*gorm.DB, error) {
	wire.Build(
		config.NewServiceConfig,
		mysql.NewGenDBClient,
	)
	return &gorm.DB{}, nil
}
