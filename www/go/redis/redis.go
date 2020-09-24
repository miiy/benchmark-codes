package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/miiy/benchmark-tools/go/config"
	"time"
)


var pool *redis.Pool

func NewPool(config *config.Redis) *redis.Pool {
	pool = redisPool(config)
	return pool
}

func NewConn(config *config.Redis) (redis.Conn, error) {
	c, err := redis.Dial("tcp", config.Host + ":" + config.Port)
	if err != nil {
		return c, err
	}
	if config.Password != "" {
		if _, err := c.Do("AUTH", config.Password); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func GetHello(c redis.Conn) string {
	hello, err := redis.String(c.Do("GET", "hello"))
	if err != nil && err != redis.ErrNil {
		fmt.Println(err)
	}
	if err == redis.ErrNil {
		fmt.Println(err)
		_, err = redis.String(c.Do("SET", "hello", "hello!", "EX", 60))
		if err != nil {
			fmt.Println(err)
		}
		hello, err = redis.String(c.Do("GET", "hello"))
	}
	return hello
}

func redisPool(config *config.Redis) *redis.Pool {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			c, err := redis.Dial("tcp", config.Host + ":" + config.Port)
			if err != nil {
				return nil, err
			}

			if config.Password != "" {
				if _, err := c.Do("AUTH", config.Password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", config.Database); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		DialContext:     nil,
		// Other pool configuration not shown in this example.
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         100,
		MaxActive:       200,
		IdleTimeout:     300 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	return pool
}

