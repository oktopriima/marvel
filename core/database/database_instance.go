/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 03/08/21, 15:57
 * Copyright (c) 2021
 */

package database

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/oktopriima/marvel/core/config"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type dbInstance struct {
	database *gorm.DB
	pool     *redis.Pool
	mgb      *mongo.Database
}

type DBInstance interface {
	Database() *gorm.DB
	Redis() *redis.Pool
	MongoDb() *mongo.Database
	Close()
}

func NewDatabaseInstance(cfg config.AppConfig) DBInstance {
	ins := new(dbInstance)

	// create connection into default database
	database, err := MysqlConnector(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed connect into database. error : %s", err.Error()))
	}
	ins.database = database

	// call redis pool
	redisPool, err := RedisConnection(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed connect into redis. error : %s", err.Error()))
	}
	ins.pool = redisPool

	// mongo db
	mongoDb, err := MongoConnection(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed connect into mongodb. error : %s", err.Error()))
	}
	ins.mgb = mongoDb

	return ins
}

func (i *dbInstance) Database() *gorm.DB {
	return i.database
}

func (i *dbInstance) Redis() *redis.Pool {
	return i.pool
}

func (i *dbInstance) MongoDb() *mongo.Database {
	return i.mgb
}

func (i *dbInstance) Close() {

}
