package controller

import (
    "encoding/json"
    "fmt"
    "github.com/go-martini/martini"
    "strconv"
    "www/app/model"
    "www/godb"
)

func Order_list(params martini.Params, db godb.Gdb) string {

    page, err := strconv.ParseUint(params["page"], 10, 64)
    if err != nil {
        fmt.Println("parse int err:", err)
    }

    var pagesize uint64
    pagesize = 15

    order := model.Order_list(page, pagesize, db)
    //model.TestRedis()
    orderInfo, err := json.Marshal(order)
    if err != nil {
        fmt.Println("order list err:", err)
    }
    return string(orderInfo)
}
