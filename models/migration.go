package models

func migration() {
	// 自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&Address{},
			&Carousels{},
			&Cart{},
			&Favorites{},
			&Orders{},
			&Products{},
			&Users{},
			&UsersAuths{},
		)
	if err != nil {
		return
	}
}
