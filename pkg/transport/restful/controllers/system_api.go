// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/services"
	"github.com/itcloudy/base-framework/pkg/transport/restful/common"
	"net/http"
)

type SystemAPIController struct {
	services.SystemAPIService
}

//根据ID获得接口详情
func (ctl SystemAPIController) CtlGetSystemAPIByID(c *gin.Context) {
	api, err := ctl.SystemAPIService.ServiceGetSystemAPIByID(c.Param("id"))
	if err != nil {

	}
	c.JSON(http.StatusOK, api)
}

//创建接口
func (ctl SystemAPIController) CtlCreateSystemAPI(c *gin.Context) {

}

//更新接口
func (ctl SystemAPIController) CtlUpdateSystemAPIByID(c *gin.Context) {

}

//查询的接口
func (ctl SystemAPIController) CtlGetAllSystemAPI(c *gin.Context) {
	apis, _, err := ctl.SystemAPIService.ServiceGetAllSystemAPI(0, consts.DefaultLimit, "", "")
	if err != nil {

	}
	common.GenResponse(c, consts.Success, apis, "")
}
