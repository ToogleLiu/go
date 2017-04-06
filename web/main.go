package main

import (
    "log"
    "net/http"
    "web/router" //导入路由
)

func main() {
    //加载路由
    router.Load()

    //监听8080端口
    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
