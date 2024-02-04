package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

// InitDB  初始化数据库
func InitDB(connString string) {
	// gorm连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               connString, // DSN data source name
		DefaultStringSize: 255,        // 默认的字符大小
	}), &gorm.Config{
		SkipDefaultTransaction: false, //为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "QI-", // 表的命名的前缀
			SingularTable: true, // 表的命名是否不带s
		},
		DisableForeignKeyConstraintWhenMigrating: true, //GORM 会自动创建外键约束,增删改会块很多
	})
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	//设置连接池
	sqlDB.SetMaxIdleConns(20)                  // 设置最大空闲连接数。
	sqlDB.SetMaxOpenConns(100)                 // 设置最大打开连接数。
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 设置连接的最大生命周期。

	DB = db

	// 数据库迁移
	migration()
}
