package main

import (
	"fmt"
	"go-chat-server/config"
	"go-chat-server/internal/infrastructure/database/mysql"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/hama?charset=utf8mb4&parseTime=True&loc=Local", config.Get().Db.User, config.Get().Db.Passwrod)
	if err := mysql.InitDb(dsn); err != nil {
		panic(fmt.Sprintf("failed to init db: %s", err.Error()))
	}
}
