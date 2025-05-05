/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 21/12/23, 16:02
 * Copyright (c) 2023
 */

package cache

import (
	"fmt"
	"github.com/oktopriima/marvel/pkg/config"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func RedisConnection(config config.AppConfig) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", config.Redis.Address, config.Redis.Port)

	idle, err := strconv.Atoi(config.Redis.MaxIdle)
	if err != nil {
		return nil, err
	}

	active, err := strconv.Atoi(config.Redis.MaxActive)
	if err != nil {
		return nil, err
	}

	db := redis.NewClient(&redis.Options{
		Addr:           addr,
		Password:       config.Redis.Password,
		DB:             idle,
		MaxActiveConns: active,
	})

	return db, nil
}

type Instance struct {
	Redis *redis.Client
}

type RedisInstance interface {
	Database() *redis.Client
	Close()
}

func NewRedisInstance(cfg config.AppConfig) RedisInstance {
	ins := new(Instance)

	// create a connection into a default database
	pool, err := RedisConnection(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to redis: %v", err))
	}
	ins.Redis = pool

	return ins
}

func (i *Instance) Database() *redis.Client {
	return i.Redis
}

func (i *Instance) Close() {
}
