/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 07/06/21, 08:41
 * Copyright (c) 2021
 */

package database

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func MysqlConnector(param map[string]interface{}) (*gorm.DB, error) {

	jsonb, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	c := new(MysqlConfig)

	if err := json.Unmarshal(jsonb, &c); err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
