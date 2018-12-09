// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package conf

import (
	"fmt"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func GetDBConnectionString(dbType string, host string, port int, user string, pass string, dbName string, charset string) (str string) {

	if dbType == "postgres" {
		str = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbName, pass)
	} else if dbType == "mysql" {

		str = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", user, pass, host, port, dbName, charset)
	}
	return
}

// GetDBConnection is initializes sqlx connection
func GetDBConnection(dbType string, host string, port int, user string, pass string, dbName string, charset string, action string) error {
	var err error
	connStr := GetDBConnectionString(dbType, host, port, user, pass, dbName, charset)
	if action != "" {
		SqlxDB, err = sqlx.Open(dbType, connStr)
	} else {
		DBConn, err = gorm.Open(dbType, connStr)
	}

	if err != nil {
		logs.Logger.Error("can't open connection to DB", zap.String("type", consts.DBError), zap.Error(err))
		DBConn = nil

		return err
	}
	return nil
}
