// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/pkg/services"
	"github.com/itcloudy/base-framework/pkg/transport/restful/common"
	"github.com/itcloudy/base-framework/tools"
	"net/http"
)

type SystemAPIController struct {
	services.SystemAPIService
}

//根据ID获得接口详情
func (ctl SystemAPIController) CtlGetSystemAPIByID(c *gin.Context) {
	api, err := ctl.SystemAPIService.ServiceGetSystemAPIByID(tools.StrToInt(c.Param("id")))
	if err != nil {

	}
	c.JSON(http.StatusOK, api)
}

//创建接口
func (ctl SystemAPIController) CtlCreateSystemAPI(c *gin.Context) {
	var (
		systemAPI models.SystemApiCreate
		result    models.SystemApiDetail
		err       error
	)
	if err = c.ShouldBindJSON(&systemAPI); err != nil {
		common.GenResponse(c, consts.BindingJsonErr, nil, "bing json failed")
		return
	}
	result, err = ctl.ServiceSystemAPICreate(systemAPI)
	if err != nil {
		common.GenResponse(c, consts.DBInSertErr, nil, err.Error())
		return
	}
	common.GenResponse(c, consts.Success, result, "")
}

type activeBind struct {
	Ids    []int `json:"ids"`
	Active bool  `json:"active"`
}

// 启用禁用接口
func (ctl SystemAPIController) CtlActiveActionSystemAPI(c *gin.Context) {
	var (
		bind activeBind
		err  error
	)
	if err = c.BindJSON(&bind); err != nil {
		common.GenResponse(c, consts.BindingJsonErr, "", err.Error())
		return
	}

	if err = ctl.ServiceActiveSystemAPI(bind.Ids, bind.Active); err != nil {
		common.GenResponse(c, consts.DBUpdateErr, "", err.Error())

	} else {
		common.GenResponse(c, consts.Success, "", "")

	}

}

//更新接口
func (ctl SystemAPIController) CtlUpdateSystemAPIByID(c *gin.Context) {

}

//查询的接口
func (ctl SystemAPIController) CtlGetAllSystemAPI(c *gin.Context) {
	var (
		page  int
		size  int
		order string
	)
	result := make(map[string]interface{})
	size = tools.StringToIntDefault(c.Query("pageSize"), consts.DefaultSize)
	page = tools.StringToIntDefault(c.Query("currentPage"), consts.DefaultPage)
	list, pagination, err := ctl.SystemAPIService.ServiceGetAllSystemAPI(page, size, order, "")
	if err != nil {

	}
	result["pagination"] = pagination
	result["list"] = list
	common.GenResponse(c, consts.Success, result, "")
}
