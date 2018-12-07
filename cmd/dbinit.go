// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package cmd

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"time"
)

var dbinitCmd = &cobra.Command{
	Use:    "dbinit",
	Short:  "init database,create tables",
	PreRun: loadConfig,
	Run: func(cmd *cobra.Command, args []string) {
		if err := models.InitDatabase(conf.Config.DB, logs.Logger); err != nil {
			logs.Logger.Fatal("init database failed", zap.String("type", conf.Config.DB.Name), zap.Error(err))

		} else {
			logs.Logger.Info("init database success", zap.String("db name", conf.Config.DB.Name), zap.Time("time", time.Now()))

		}

	},
}
