/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 21/12/23, 16:02
 * Copyright (c) 2023
 */

package database

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/oktopriima/marvel/pkg/config"
	"strconv"
)

func RedisConnection(config config.AppConfig) (*redis.Pool, error) {
	addr := fmt.Sprintf("%s:%s", config.Redis.Address, config.Redis.Port)

	idle, err := strconv.Atoi(config.Redis.MaxIdle)
	if err != nil {
		return nil, err
	}

	active, err := strconv.Atoi(config.Redis.MaxActive)
	if err != nil {
		return nil, err
	}

	return &redis.Pool{
		MaxIdle:   idle,
		MaxActive: active,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH", config.Redis.Password); err != nil {
				return nil, err
			}

			return c, nil
		},
	}, nil
}
