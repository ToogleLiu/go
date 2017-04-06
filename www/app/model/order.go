package model

import (
    "fmt"
    "www/godb"
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

func Order_list(page uint64, pagesize uint64, db godb.Gdb) OrderSlice {

    start := (page - 1) * pagesize

    var odb = db.DB_order

    //db.Where("order_id > ?", 0).Order("order_id desc").Limit(15).Offset(0).Find(&order)

    rows, err := odb.Model(&Order{}).Where("order_id > ?", 0).Order("order_id desc").Limit(pagesize).Offset(start).Rows()
    defer rows.Close()

    if err != nil {
        panic("query failed.")
    }

    var orders OrderSlice

    for rows.Next() {
        var order Order
        odb.ScanRows(rows, &order)
        orders.Orders = append(orders.Orders, order)
    }

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
