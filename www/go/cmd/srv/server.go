package main

import (
	goredis "github.com/gomodule/redigo/redis"
	"github.com/miiy/benchmark-tools/go/config"
	"github.com/miiy/benchmark-tools/go/redis"
	"net/http"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!"))
	})
	http.HandleFunc("/redis", redisHandler)
	http.HandleFunc("/redis-pool", redisPoolHandler)

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func redisHandler(w http.ResponseWriter, r *http.Request)  {
	c, err := redis.NewConn(&conf.Redis)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	hello := redis.GetHello(c)
	w.Write([]byte(hello))
}

func redisPoolHandler(w http.ResponseWriter, r *http.Request) {
	c := pool.Get()

	hello := redis.GetHello(c)
	w.Write([]byte(hello))
}