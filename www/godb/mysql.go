package godb

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "www/config"
)

type Gdb struct {
    DB_order *gorm.DB
    DB_user  *gorm.DB
}

func open(config config.DBConfig) *gorm.DB {
    openInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DBName)
    //db, err := gorm.Open("mysql", "N5lux:idbT3W3n0@/v30_db_orders?charset=utf8&parseTime=True&loc=Local")
    db, err := gorm.Open("mysql", openInfo)
    if err != nil {
        panic("failed to connect database")
    }
    //  defer db.Close()
    return db
}

func OpenDB() Gdb {
    var db Gdb
    db.DB_order = open(config.DB_order())
    db.DB_user = open(config.DB_user())
    return db
}

func CloseDB(db Gdb) {
    db_order := db.DB_order
    db_user := db.DB_user
    db_order.Close()
    db_user.Close()
}
