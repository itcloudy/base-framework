// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/services"
	"net/http"
)

type MenuController struct {
	services.MenuService
}

// 用户获取自己的菜单
func (ctl MenuController) CtlGetSelfMenu(c *gin.Context) {

	menus, err := ctl.MenuService.ServiceGetSelfMenu(c.GetStringSlice(consts.LoginUserRoleIds))
	if err != nil {

	}
	c.JSON(http.StatusOK, menus)
}
//根据ID获得菜单详情
func (ctl MenuController)CtlGetMenuByID(c*gin.Context)  {

}
//创建菜单
func (ctl MenuController)CtlCreateMenu(c*gin.Context)  {

}
//更新菜单
func (ctl MenuController)CtlUpdateMenuByID(c*gin.Context)  {

}
//获得所有的菜单
func (ctl MenuController)CtlGetAllMenu(c*gin.Context)  {

}