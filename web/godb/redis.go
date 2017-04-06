package godb

import (
    "fmt"
    "gopkg.in/redis.v4"
)

func RedisOpen() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "94f20352-a9eb-4dfd-adfb-bdc2dd9b03a4:deen112255", //deen112255
        DB:       0,
    })

    // 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
    pong, err := client.Ping().Result()
    fmt.Println(pong, err)

    return client
}
