package main

import (
	"github.com/gin-gonic/gin"
	goredis "github.com/gomodule/redigo/redis"
	"github.com/miiy/benchmark-tools/go/config"
	"github.com/miiy/benchmark-tools/go/redis"
)

var conf *config.Config

var pool *goredis.Pool

func main()  {

	conf = &config.Config{
		Redis: config.Redis{
			Host: "127.0.0.1",
			Port: "6379",
			Password: "",
			Database: 0,
		},
	}
	pool = redis.NewPool(&conf.Redis)


	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("hello!"))
	})
	r.GET("/redis", redisHandler)
	r.GET("/redis-pool", redisPoolHandler)
	r.Run("127.0.0.1:8080")
}


func redisHandler(ctx *gin.Context) {
	c, err := redis.NewConn(&conf.Redis)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	hello := redis.GetHello(c)
	ctx.Writer.Write([]byte(hello))
}

func redisPoolHandler(ctx *gin.Context) {
	c := pool.Get()
	defer c.Close()

	hello := redis.GetHello(c)
	ctx.Writer.Write([]byte(hello))
}