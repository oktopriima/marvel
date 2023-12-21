/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 21/12/23, 16:02
 * Copyright (c) 2023
 */

package database

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type RedisClient struct {
	Address   string `json:"address"`
	Port      string `json:"port"`
	Password  string `json:"password"`
	MaxIdle   string `json:"max_idle"`
	MaxActive string `json:"max_active"`
}

func RedisConnection(c map[string]interface{}) (*redis.Pool, error) {
	var client RedisClient

	jsonb, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(jsonb, &client); err != nil {
		return nil, fmt.Errorf("unable to parse config: %v", err)
	}

	addr := fmt.Sprintf("%s:%s", client.Address, client.Port)

	idle, err := strconv.Atoi(client.MaxIdle)
	if err != nil {
		return nil, err
	}

	active, err := strconv.Atoi(client.MaxActive)
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

			if _, err := c.Do("AUTH", client.Password); err != nil {
				return nil, err
			}

			return c, nil
		},
	}, nil
}
