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

// @tags  菜单
// @Description 用户菜单获取
// @Summary 用户菜单获取
// @Accept  json
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"token":"Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjc1MzIxNjMwODMsIk5hbWUiOiJhZG1pbiIsIlJvbGUiOm51bGwsIlVzZXJJZCI6MiwiSXNBZG1pbiI6ZmFsc2V9.HZq5jBw4-ZQipQPnq0K7Ei0_LvaRXZGNgKqLoFnhV_vpfQupmddsDMZbiI_Yy0Zhd7J7AvRGDXMfVwW9-TidsDrux6-L4KQWIV0Mrlj4SXgW13HvMSXW0XzHYQBxiai61AeJx4VmQR84s2lI5hmKuiVOpsyOZAduJoO1K26b8X4","user":{"id":2,"name":"admin","alias":"","email":"","password":"","roles":[],"openid":"admin","active":true,"is_admin":false}},"message":"success"}"
// @Router /user/{id} [get]
func (ctl *MenuController) CtlGetSelfMenu(c *gin.Context) {

	menus, err := ctl.MenuService.GetSelfMenu(c.GetStringSlice(consts.LoginUserRoles))
	if err != nil {

	}
	c.JSON(http.StatusOK, menus)
}
