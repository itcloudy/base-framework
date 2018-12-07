// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package cmd

import "github.com/spf13/cobra"

var loaddataCmd = &cobra.Command{
	Use:    "loaddata",
	Short:  "load data from init folder",
	PreRun: loadConfig,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
