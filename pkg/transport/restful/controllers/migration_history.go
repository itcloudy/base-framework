// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/interfaces/services"
	"github.com/itcloudy/base-framework/pkg/transport/restful/common"
)

type MigrationHistoryController struct {
	services.IMigrationHistory
}

func (ctl MigrationHistoryController) CtlGetAllMigrations(c *gin.Context) {
	results, err := ctl.ServiceGetAllListMigration()
	if err != nil {
		common.GenResponse(c, consts.DBSelectErr, "", err.Error())
	} else {
		common.GenResponse(c, consts.Success, results, "")
	}
}

func (ctl MigrationHistoryController) CtlApplyMigration(c *gin.Context) {
	version := c.Param("version")
	if len(version) == 0 {
		common.GenResponse(c, consts.PathParamErr, "", "version is empty")
	}
	err := ctl.ServiceUpdateToOneVersion(version)
	if err != nil {
		common.GenResponse(c, consts.MigrationErr, "", err.Error())

	} else {
		common.GenResponse(c, consts.Success, "", "")

	}
}
