package controller

import (
    "fmt"
    "strconv"
    "net/http"
//  "strings"
    "encoding/json"
//  "web/godb"
    "web/app/model"
)


func Order_list(w http.ResponseWriter, r *http.Request) {

    r.ParseForm()
    page, err := strconv.ParseUint(r.Form["page"][0], 10, 64)
    if err != nil {
        fmt.Println("parse int err:", err)
    }

    var pagesize uint64
    pagesize = 15

    order := model.Order_list(page, pagesize)
//  model.TestRedis()
    orderInfo, err := json.Marshal(order)
    if err != nil {
        fmt.Println("order list err:", err)
    }
    fmt.Fprintf(w, string(orderInfo))   
}
