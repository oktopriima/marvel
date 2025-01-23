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
	"github.com/oktopriima/marvel/pkg/config"
	"gorm.io/gorm"
)

type Instance struct {
	GormDB *gorm.DB
}

type DBInstance interface {
	Database() *gorm.DB
	Close()
}

func NewDatabaseInstance(cfg config.AppConfig) DBInstance {
	ins := new(Instance)

	// create connection into default database
	database, err := MysqlConnector(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed connect into database. error : %s", err.Error()))
	}
	ins.GormDB = database

	return ins
}

func (i *Instance) Database() *gorm.DB {
	return i.GormDB
}

func (i *Instance) Close() {

}
