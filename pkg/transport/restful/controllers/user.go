// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/interfaces/services"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/pkg/transport/restful/common"
	"github.com/itcloudy/base-framework/pkg/transport/restful/middles"
	"github.com/itcloudy/base-framework/tools"
	"net/http"
)

type UserController struct {
	services.IUserService
}

// @tags  用户
// @Description 根据用户ID获取用户信息
// @Summary 用户信息获取
// @Accept  json
// @Produce  json
// @Param id path string true "用户ID"
// @Success 200 {string} json "{"code":200,"data":{"token":"Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjc1MzIxNjMwODMsIk5hbWUiOiJhZG1pbiIsIlJvbGUiOm51bGwsIlVzZXJJZCI6MiwiSXNBZG1pbiI6ZmFsc2V9.HZq5jBw4-ZQipQPnq0K7Ei0_LvaRXZGNgKqLoFnhV_vpfQupmddsDMZbiI_Yy0Zhd7J7AvRGDXMfVwW9-TidsDrux6-L4KQWIV0Mrlj4SXgW13HvMSXW0XzHYQBxiai61AeJx4VmQR84s2lI5hmKuiVOpsyOZAduJoO1K26b8X4","user":{"id":2,"name":"admin","alias":"","email":"","password":"","roles":[],"openid":"admin","active":true,"is_admin":false}},"message":"success"}"
// @Router /user/{id} [get]
func (ctl UserController) CtlGetUserByID(c *gin.Context) {
	user, err := ctl.ServiceGetUserByID(tools.StrToInt(c.Param("id")))
	if err != nil {
	}
	c.JSON(http.StatusOK, user)
}

// @tags  用户
// @Description 根据用户名获取用户信息
// @Summary 用户信息获取
// @Accept  json
// @Produce  json
// @Param username path string true "用户名"
// @Success 200 {string} json "{"code":200,"data":{"token":"Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjc1MzIxNjMwODMsIk5hbWUiOiJhZG1pbiIsIlJvbGUiOm51bGwsIlVzZXJJZCI6MiwiSXNBZG1pbiI6ZmFsc2V9.HZq5jBw4-ZQipQPnq0K7Ei0_LvaRXZGNgKqLoFnhV_vpfQupmddsDMZbiI_Yy0Zhd7J7AvRGDXMfVwW9-TidsDrux6-L4KQWIV0Mrlj4SXgW13HvMSXW0XzHYQBxiai61AeJx4VmQR84s2lI5hmKuiVOpsyOZAduJoO1K26b8X4","user":{"id":2,"name":"admin","alias":"","email":"","password":"","roles":[],"openid":"admin","active":true,"is_admin":false}},"message":"success"}"
// @Router /user/{id} [get]
func (ctl UserController) CtlGetUserByUserName(c *gin.Context) {
	user, err := ctl.ServiceGetUserByUserName(c.Param("username"))
	if err != nil {

	}
	common.GenResponse(c, consts.Success, user, "")
}

// 获得登录用户的信息
func (ctl UserController) CtlGetSelfInformation(c *gin.Context) {
	user, err := ctl.ServiceGetUserByID(c.GetInt(consts.LoginUserID))
	if err != nil {
		common.GenResponse(c, consts.ServerErr, "", "")
		return
	}
	response := make(map[string]interface{})
	response["information"] = user
	common.GenResponse(c, consts.Success, response, "")
}
func (ctl UserController) CtlRefreshToken(c *gin.Context) {

}
func (ctl UserController) CtlLoginAccount(c *gin.Context) {
	var (
		user       models.UserAuth
		userDetail models.UserDetail
		err        error
	)
	if err = c.ShouldBindJSON(&user); err != nil {
		common.GenResponse(c, consts.BindingJsonErr, nil, "bing json failed")
		return
	}
	response := make(map[string]interface{})
	response["type"] = "account"
	code := consts.UserNameOrPasswordErr
	if userDetail, err = ctl.ServiceCheckUser(user.Username, user.Password); err != nil || userDetail.ID == 0 {
		response["currentAuthority"] = "guest"
	} else {
		code = consts.Success
		response["currentAuthority"] = "admin"
		response["token"] = middles.GenerateJWT(userDetail.Username,
			[]string{}, []string{}, userDetail.ID, userDetail.IsAdmin)
	}
	common.GenResponse(c, code, response, "")
}

//查询的角色
func (ctl UserController) CtlGetAllUser(c *gin.Context) {
	var (
		page  int
		size  int
		order string
	)
	size = tools.StringToIntDefault(c.Query("pageSize"), consts.DefaultSize)
	page = tools.StringToIntDefault(c.Query("currentPage"), consts.DefaultPage)
	users, _, err := ctl.ServiceGetAllUser(page, size, order, "")
	if err != nil {

	}
	common.GenResponse(c, consts.Success, users, "")
}
