// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package cmd

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/pkg/models"
	"go.uber.org/zap"
	"time"

	"github.com/spf13/cobra"
)

var dbupdateCmd = &cobra.Command{
	Use:    "dbupdate",
	Short:  "update database, add tables and columns or modify columns",
	PreRun: loadConfig,
	Run: func(cmd *cobra.Command, args []string) {
		if err := models.UpdateDatabase(conf.Config.DB); err != nil {
			logs.Logger.Fatal("update database failed", zap.String("db name", conf.Config.DB.Name), zap.Error(err))
			logs.Logger.Sync()

		} else {
			logs.Logger.Info("update database success", zap.String("db name", conf.Config.DB.Name), zap.Time("time", time.Now()))
			logs.Logger.Sync()

		}
	},
}
