package model

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "web/godb"
)

type Order struct {
    Order_id uint   `gorm:"primary_key;AUTO_INCREMENT" json:"order_id"`
    Order_sn string `json:"order_sn"`
    Add_time uint   `json:"add_time"`
}

type OrderSlice struct {
    Orders []Order `json:"orders"`
}

func (Order) TableName() string {
    return "ecs_order_info"
}

func openDB() *gorm.DB {
    return godb.Open(godb.DB_order())
}

func closeDB(db *gorm.DB) {
    godb.Close(db)
}

func Order_list(page uint64, pagesize uint64) OrderSlice {
    db := openDB()

    //var order Order

    start := (page - 1) * pagesize

    //db.Where("order_id > ?", 0).Order("order_id desc").Limit(15).Offset(0).Find(&order)
    rows, err := db.Model(&Order{}).Where("order_id > ?", 0).Order("order_id desc").Limit(pagesize).Offset(start).Rows()
    defer rows.Close()

    if err != nil {
        panic("query failed.")
    }

    var orders OrderSlice

    for rows.Next() {
        var order Order
        db.ScanRows(rows, &order)
        orders.Orders = append(orders.Orders, order)
    }

    //db.Last(&order)

    closeDB(db)

    return orders
}

func TestRedis() {
    redis := godb.RedisOpen()
    err := redis.Set("goredis_name", "hello,go.from redis.", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := redis.Get("goredis_name").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("name", val)
}
