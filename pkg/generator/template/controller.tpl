// Copyright 2018  itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"{{.ProjectPath}}/pkg/consts"
	"{{.ProjectPath}}/pkg/interfaces/services"
	"{{.ProjectPath}}/pkg/models"
	"{{.ProjectPath}}/pkg/restful/common"
	"{{.ProjectPath}}/tools"
	"github.com/gin-gonic/gin"
)

type {{.ModelName}}Controller struct {
	services.I{{.ModelName}}Service
}

//根据ID获得详情
func (ctl {{.ModelName}}Controller) CtlGet{{.ModelName}}ByID(c *gin.Context) {
	result, err := ctl.ServiceGet{{.ModelName}}ByID(tools.StrToInt(c.Param("id")))
	if err != nil {
        common.GenResponse(c, consts.DBSelectErr, "", err.Error())
	}
	common.GenResponse(c, consts.Success, result, "")
}

//创建
func (ctl {{.ModelName}}Controller) CtlCreate{{.ModelName}}(c *gin.Context) {
	var (
		create models.{{.ModelName}}Create
		result    models.{{.ModelName}}Detail
		err       error
	)
	if err = c.ShouldBindJSON(&create); err != nil {
		common.GenResponse(c, consts.BindingJsonErr, nil, "bind json failed")
		return
	}
	result, err = ctl.Service{{.ModelName}}Create(create)
	if err != nil {
		common.GenResponse(c, consts.DBInSertErr, nil, err.Error())
		return
	}
	common.GenResponse(c, consts.Success, result, "")
}

//更新
func (ctl {{.ModelName}}Controller) CtlUpdate{{.ModelName}}ByID(c *gin.Context) {
	var (
		model models.{{.ModelName}}Update
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
	result, err := ctl.Service{{.ModelName}}Update(model)
	if err != nil {
		common.GenResponse(c, consts.PathParamErr, "", err.Error())
	} else {
		common.GenResponse(c, consts.Success, result, "")
	}
}
//删除
func (ctl {{.ModelName}}Controller) Ctl{{.ModelName}}Delete(c *gin.Context) {
	var (
		bind activeBind
		err  error
	)
	if err = c.BindJSON(&bind); err != nil {
		common.GenResponse(c, consts.BindingJsonErr, "", err.Error())
		return
	}

	if err = ctl.Service{{.ModelName}}Delete(bind.Ids); err != nil {
		common.GenResponse(c, consts.DBDeleteErr, "", err.Error())

	} else {
		common.GenResponse(c, consts.Success, "", "")
	}

}
//查询
func (ctl {{.ModelName}}Controller) CtlGetAll{{.ModelName}}(c *gin.Context) {
	var (
		page  int
		size  int
		order string
	)
	result := make(map[string]interface{})
	size = tools.StringToIntDefault(c.Query("pageSize"), consts.DefaultSize)
	page = tools.StringToIntDefault(c.Query("currentPage"), consts.DefaultPage)
	list, pagination, err := ctl.ServiceGetAll{{.ModelName}}(page, size, order, "")
	if err != nil {

	}
	result["pagination"] = pagination
	result["list"] = list
	common.GenResponse(c, consts.Success, result, "")
}

