package main

import (
	"github.com/soladxy/oyasumi/biz/container"
	"gorm.io/gen"
)

func main() {
	db, err := container.InitGormGen("./conf/oyasumi.yaml")
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		// 最终package不能设置为model，在有数据库表同步的情况下会产生冲突，若一定要使用可以单独指定model package的新名字
		OutPath:      "./biz/dal/mysql/query",
		ModelPkgPath: "./biz/dal/mysql/model", // 默认情况下会跟随OutPath参数，在同目录下生成model目录
	})

	g.UseDB(db) // reuse your gorm db

	// generate all table from database
	//g.ApplyBasic(g.GenerateAllTable()...)

	g.ApplyBasic(g.GenerateModel("website"))
	// Generate the code
	g.Execute()
}
