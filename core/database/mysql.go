/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 07/06/21, 08:41
 * Copyright (c) 2021
 */

package database

import (
	"errors"
	"fmt"
	"github.com/oktopriima/marvel/core/config"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"io/fs"
	"log"
	"os"
	"time"
)

func MysqlConnector(cfg config.AppConfig) (*gorm.DB, error) {
	var dbLogFile *os.File
	dbLogFile, err := os.OpenFile(fmt.Sprintf("%s/%s", cfg.Log.Directory, cfg.Log.Mysql), os.O_CREATE|os.O_RDWR|os.O_APPEND, fs.ModeAppend)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		_ = os.Mkdir(cfg.Log.Directory, os.ModePerm)
		dbLogFile, err = os.Create(fmt.Sprintf("%s/%s", cfg.Log.Directory, cfg.Log.Mysql))
	}

	dbLogger := logger.New(
		log.New(io.MultiWriter(os.Stdout, dbLogFile), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	gormConfig := &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 dbLogger,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.Database,
	)

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

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
