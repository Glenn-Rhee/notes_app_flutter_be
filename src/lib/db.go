package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnect() (*gorm.DB, error) {
	// DSN (Data Source Name) untuk koneksi ke database MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/notes_app_flutter?charset=utf8mb4&parseTime=True&loc=Local"

	// Membuka koneksi ke database menggunakan GORM dengan driver MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return  nil, err
	}
	
	return db, nil
}