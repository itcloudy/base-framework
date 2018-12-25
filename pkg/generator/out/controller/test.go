// Copyright 2018 itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/interfaces/services"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/pkg/restful/common"
	"github.com/itcloudy/base-framework/tools"
	"net/http"
)

type TestController struct {
	services.ITestService
}

//根据ID获得详情
func (ctl TestController) CtlGetTestByID(c *gin.Context) {
	api, err := ctl.ServiceGetTestByID(tools.StrToInt(c.Param("id")))
	if err != nil {

	}
	c.JSON(http.StatusOK, api)
}

//创建
func (ctl TestController) CtlCreateTest(c *gin.Context) {
	var (
		create models.TestCreate
		result models.TestDetail
		err    error
	)
	if err = c.ShouldBindJSON(&create); err != nil {
		common.GenResponse(c, consts.BindingJsonErr, nil, "bind json failed")
		return
	}
	result, err = ctl.ServiceTestCreate(create)
	if err != nil {
		common.GenResponse(c, consts.DBInSertErr, nil, err.Error())
		return
	}
	common.GenResponse(c, consts.Success, result, "")
}

//更新
func (ctl TestController) CtlUpdateTestByID(c *gin.Context) {
	var (
		model models.TestUpdate
		err   error
	)
	id := tools.StrToInt(c.Param("id"))
	if id < 1 {
		common.GenResponse(c, consts.UpdateIdErr, "", "")
	}
	if err = c.BindJSON(&model); err != nil {
		common.GenResponse(c, consts.BindingJsonErr, "", err.Error())
		return
	}
	model.ID = id
	result, err := ctl.ServiceTestUpdate(model)
	if err != nil {
		common.GenResponse(c, consts.PathParamErr, "", err.Error())
	} else {
		common.GenResponse(c, consts.Success, result, "")
	}
}
func (ctl TestController) CtlDeleteTest(c *gin.Context) {
	var (
		bind activeBind
		err  error
	)
	if err = c.BindJSON(&bind); err != nil {
		common.GenResponse(c, consts.BindingJsonErr, "", err.Error())
		return
	}

	if err = ctl.ServiceDeleteTest(bind.Ids); err != nil {
		common.GenResponse(c, consts.DBDeleteErr, "", err.Error())

	} else {
		common.GenResponse(c, consts.Success, "", "")
	}

}

//查询
func (ctl TestController) CtlGetAllTest(c *gin.Context) {
	var (
		page  int
		size  int
		order string
	)
	result := make(map[string]interface{})
	size = tools.StringToIntDefault(c.Query("pageSize"), consts.DefaultSize)
	page = tools.StringToIntDefault(c.Query("currentPage"), consts.DefaultPage)
	list, pagination, err := ctl.ServiceGetAllTest(page, size, order, "")
	if err != nil {

	}
	result["pagination"] = pagination
	result["list"] = list
	common.GenResponse(c, consts.Success, result, "")
}
