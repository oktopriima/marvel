/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 13/08/21, 14:31
 * Copyright (c) 2021
 */

package nosql

import (
	"context"
	"fmt"
	"github.com/oktopriima/marvel/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection(config config.AppConfig) (*mongo.Database, error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.Mongodb.User, config.Mongodb.Password, config.Mongodb.Address, config.Mongodb.Port)
	clientOptions := options.Client().ApplyURI(url)

	ctx := context.Background()
	mgo, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	if err = mgo.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	db := mgo.Database(config.Mongodb.Database)
	return db, nil
}

type Instance struct {
	Mongodb *mongo.Database
}

type MongoInstance interface {
	Database() *mongo.Database
	Close()
}

func NewMongoDBInstance(cfg config.AppConfig) MongoInstance {
	ins := new(Instance)

	// create a connection into a default database
	mongodb, err := MongoConnection(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to redis: %v", err))
	}
	ins.Mongodb = mongodb

	return ins
}

func (i *Instance) Database() *mongo.Database {
	return i.Mongodb
}

func (i *Instance) Close() {
}
