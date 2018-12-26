// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package cmd

import (
	"fmt"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/logs"

	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "base-framework",
	Short: "base-framework application",
}

func init() {
	rootCmd.AddCommand(
		configCmd,
		dbinitCmd,
		dbupdateCmd,
		loaddataCmd,
		startCmd,
	)

	// This flags are visible for all child commands
	rootCmd.PersistentFlags().StringVar(&conf.Config.ConfigPath, "config", defaultConfigPath(), "filepath to config.toml")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Executing root command")
	}
}

func defaultConfigPath() string {
	p, err := os.Getwd()
	if err != nil {
		logs.Logger.Fatal("getting current wd")
	}
	return filepath.Join(p, "config", "config.toml")
}

// Load the configuration from file
func loadConfig(cmd *cobra.Command, args []string) {
	err := conf.LoadConfig(conf.Config.ConfigPath)
	if err != nil {
		fmt.Printf("Loading config failed, config path: %s , err info: %+v\n", conf.Config.ConfigPath, err)
	} else {
		logConf := conf.Config.Log
		logs.InitLogger(
			conf.Config.Mode,
			logConf.FileName,
			logConf.MaxSize,
			logConf.MaxBackups,
			logConf.MaxAge,
			logConf.Compress,
			logConf.EnableKafka,
			logConf.KafkaAddress)
	}
}
