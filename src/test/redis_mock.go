package test

import (
	"github.com/go-redis/redismock/v9"
	"github.com/oktopriima/marvel/pkg/cache"
)

func RedisInstance() (cache.RedisInstance, redismock.ClientMock) {
	db, mock := redismock.NewClientMock()

	ins := new(cache.Instance)
	ins.Redis = db
	
	return ins, mock
}
