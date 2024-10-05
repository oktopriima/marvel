/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 07/06/21, 08:41
 * Copyright (c) 2021
 */

package database

import (
	"fmt"
	"github.com/oktopriima/marvel/core/config"
	"go.elastic.co/apm/module/apmsql/v2"
	_ "go.elastic.co/apm/module/apmsql/v2/mysql"
	"go.elastic.co/apm/v2"
	gormSql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func MysqlConnector(cfg config.AppConfig) (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.Database,
	)

	tracer := apm.DefaultTracer()
	if tracer == nil {
		panic("tracer is nil")
	}

	sqlMaster, err := apmsql.Open("mysql", dsn)
	if err != nil {
		panic("error during construct APM, please check the config")
	}

	// Set Max Idle Connections sets the maximum number of connections in the idle connection pool.
	sqlMaster.SetMaxIdleConns(10)
	// Set Max Open Connections sets the maximum number of open connections to the database.
	sqlMaster.SetMaxOpenConns(100)
	// Set Connections Max Lifetime sets the maximum amount of time a connection may be reused.
	sqlMaster.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(gormSql.New(gormSql.Config{Conn: sqlMaster}), gormConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
