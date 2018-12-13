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
func (ctl MenuController) CtlGetMenuByID(c *gin.Context) {

}

//创建菜单
func (ctl MenuController) CtlCreateMenu(c *gin.Context) {

}

//更新菜单
func (ctl MenuController) CtlUpdateMenuByID(c *gin.Context) {

}

//查询的菜单
func (ctl MenuController) CtlGetAllMenu(c *gin.Context) {
	var (
		page  int
		size  int
		order string
	)
	result := make(map[string]interface{})
	size = tools.StringToIntDefault(c.Query("size"), consts.DefaultSize)
	page = tools.StringToIntDefault(c.Query("current"), consts.DefaultPage)
	menus, pagination, err := ctl.MenuService.ServiceGetAllMenu(page, size, order, "")
	if err != nil {

	}
	result["pagination"] = pagination
	result["list"] = menus
	common.GenResponse(c, consts.Success, result, "")
}
