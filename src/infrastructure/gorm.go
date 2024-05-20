package infrastructure

import (
	"fmt"
	"os"

	"github.com/KotaTanaka/echo-api-sandbox/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)

	fmt.Println(dsn)

	gormDB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	gormDB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	gormDB = MigrateDB(gormDB)

	return gormDB, nil
}

func MigrateDB(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.Area{})
	db.AutoMigrate(&model.Service{})
	db.AutoMigrate(&model.Shop{})
	db.AutoMigrate(&model.Review{})

	return db
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
