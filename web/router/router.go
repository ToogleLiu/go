package router

import (
    "net/http"
    "web/app/controller" //导入应用包
)

func Load() {
    //设置路由

    //http.HandleFunc("/", controller.SayhelloName)

    http.HandleFunc("/order", controller.Order_list)

//  http.HandleFunc("/good/index", controller.Good)
}
