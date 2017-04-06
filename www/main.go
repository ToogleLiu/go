package main

import (
    //  "fmt"
    "github.com/go-martini/martini"
    "www/godb"
    "www/router"
)

func main() {
    db := godb.OpenDB()

    m := martini.Classic()

    m.Map(db)

    //加载路由
    router.Load(m, db)

    m.RunOnAddr(":3030")

    m.Run()

    godb.CloseDB(db)
}
