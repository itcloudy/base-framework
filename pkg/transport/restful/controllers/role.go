// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/services"
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

//获得所有的角色
func (ctl RoleController) CtlGetAllRole(c *gin.Context) {

}
