package init

import (
	"github.com/joho/godotenv"
	"soladxy/oyasumi/biz/dal/mysql"
)

func Init() {
	initEnv()
	mysql.InitDb()
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}
}
