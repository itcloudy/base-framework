// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package cmd

import (
	"errors"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/pkg/repositories/common"
	"github.com/itcloudy/base-framework/pkg/services"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"time"
)

var dbupdateCmd = &cobra.Command{
	Use:    "dbupdate",
	Short:  "update database, add tables and columns or modify columns",
	PreRun: loadConfig,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := conf.Config.DB
		conf.GetDBConnection(cfg.DbType, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.Charset, "update")
		migrateService := services.MigrationService{}
		migrateService.IMigrationHistoryRepository = &common.MigrationHistoryRepository{}

		switch cfg.DbType {
		case "mysql":
		case "postgres":
			migrateService.DB = conf.SqlxDB
			break
		default:
			panic(errors.New("un support sql type:" + cfg.DbType))

		}
		if err := migrateService.ServiceUpdateToOneVersion(conf.Config.DBUpdateToVersion); err != nil {
			logs.Logger.Fatal("update database failed", zap.String("db name", conf.Config.DB.Name), zap.Error(err))
			logs.Logger.Sync()

		} else {
			logs.Logger.Info("update database success", zap.String("db name", conf.Config.DB.Name), zap.Time("time", time.Now()))
			logs.Logger.Sync()
		}
		conf.SqlxDB.Close()

	},
}

func init() {
	dbupdateCmd.Flags().StringVar(&conf.Config.DBUpdateToVersion, "to-version", "", "update database to version")
	dbupdateCmd.MarkFlagRequired("to-version")
}
