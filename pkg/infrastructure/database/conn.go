package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"sample/pkg/domain/model"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// Driver名
const driverName = "mysql"

var DB *gorm.DB

func ConnectDB() {
	/* ===== データベースへ接続する. ===== */
	// ユーザ
	user := os.Getenv("MYSQL_USER")
	// パスワード
	password := os.Getenv("MYSQL_PASSWORD")
	// 接続先ホスト
	host := os.Getenv("MYSQL_HOST")
	// 接続先ポート
	port := os.Getenv("MYSQL_PORT")
	// 接続先データベース
	database := os.Getenv("MYSQL_DATABASE")

	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:        fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", user, password, host, port, database),
		DriverName: driverName,
	}))

	DB.Logger = DB.Logger.LogMode(logger.Info)

	if err != nil {
		log.Fatal(err)
	}
}

func AutoMigrate() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}
}
