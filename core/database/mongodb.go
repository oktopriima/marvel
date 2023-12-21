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
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func MongoConnection(c map[string]interface{}) (*mongo.Database, error) {
	var client MongoClient

	jsonb, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(jsonb, &client); err != nil {
		return nil, fmt.Errorf("unable to parse config: %v", err)
	}

	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", client.User, client.Password, client.Address, client.Port)
	clientOptions := options.Client().ApplyURI(url)

	mgo, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	if err = mgo.Connect(context.Background()); err != nil {
		return nil, err
	}

	if err = mgo.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	db := mgo.Database(client.Database)
	return db, nil
}
