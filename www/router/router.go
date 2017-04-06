package router

import (
    //"fmt"
    "github.com/go-martini/martini"
    "www/app/controller" //导入应用包
    "www/godb"
)

func Load(m *martini.ClassicMartini, db godb.Gdb) {
    //设置路由

    //首页
    m.Get("/", func() string {
        return "<div style=\"width:100%;height:800px;line-height:800px;font-size:30px;text-align:center;\">Hello,Go lang.</div>"
    })

    //订单列表
    m.Get("/order/:page", func(params martini.Params) string {
        return controller.Order_list(params, db)
    })
}
