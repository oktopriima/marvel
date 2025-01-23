/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 13/08/21, 14:31
 * Copyright (c) 2021
 */

package database

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
