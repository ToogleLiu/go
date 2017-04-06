package godb

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Open(config DBConfig) *gorm.DB {
    openInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DBName)
    //db, err := gorm.Open("mysql", "N5lux:idbT3W3n0@/v30_db_orders?charset=utf8&parseTime=True&loc=Local")
    db, err := gorm.Open("mysql", openInfo)
    if err != nil {
        panic("failed to connect database")
    }
    return db
}

func Close(db *gorm.DB) {
    db.Close()
}