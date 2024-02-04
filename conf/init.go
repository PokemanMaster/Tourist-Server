package conf

import (
	"github.com/joho/godotenv"
	"os"
	"qiqi-go/cache"
	"qiqi-go/models"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	// 数据库的连接
	models.InitDB(os.Getenv("MYSQL_DSN"))
	cache.InitRedis()
}
