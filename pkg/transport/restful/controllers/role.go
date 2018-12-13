// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/services"
	"github.com/itcloudy/base-framework/pkg/transport/restful/common"
	"github.com/itcloudy/base-framework/tools"
)

type RoleController struct {
	services.RoleService
}

//根据ID获得角色详情
func (ctl RoleController) CtlGetRoleByID(c *gin.Context) {

}

//创建角色
func (ctl RoleController) CtlCreateRole(c *gin.Context) {

}

//更新角色
func (ctl RoleController) CtlUpdateRoleByID(c *gin.Context) {

}

//查询的角色
func (ctl RoleController) CtlGetAllRole(c *gin.Context) {
	var (
		page  int
		size  int
		order string
	)
	size = tools.StringToIntDefault(c.Query("pageSize"), consts.DefaultSize)
	page = tools.StringToIntDefault(c.Query("currentPage"), consts.DefaultPage)
	apis, _, err := ctl.RoleService.ServiceGetAllRole(page, size, order, "")
	if err != nil {

	}
	common.GenResponse(c, consts.Success, apis, "")
}
