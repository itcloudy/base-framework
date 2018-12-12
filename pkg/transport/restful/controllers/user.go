// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
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
	user, err := ctl.ServiceGetUserByID(c.Param("id"))
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

// @tags  用户
// @Description 获得登录用户的信息
// @Summary 用户信息获取
// @Accept  json
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"token":"Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjc1MzIxNjMwODMsIk5hbWUiOiJhZG1pbiIsIlJvbGUiOm51bGwsIlVzZXJJZCI6MiwiSXNBZG1pbiI6ZmFsc2V9.HZq5jBw4-ZQipQPnq0K7Ei0_LvaRXZGNgKqLoFnhV_vpfQupmddsDMZbiI_Yy0Zhd7J7AvRGDXMfVwW9-TidsDrux6-L4KQWIV0Mrlj4SXgW13HvMSXW0XzHYQBxiai61AeJx4VmQR84s2lI5hmKuiVOpsyOZAduJoO1K26b8X4","user":{"id":2,"name":"admin","alias":"","email":"","password":"","roles":[],"openid":"admin","active":true,"is_admin":false}},"message":"success"}"
// @Router /user/{id} [get]
func (ctl UserController) CtlGetSelf(c *gin.Context) {
	user, err := ctl.ServiceGetUserByID(c.GetString(consts.LoginUserID))
	if err != nil {

	}
	common.GenResponse(c, consts.Success, user, "")
}

func (ctl UserController) CtlLogin(c *gin.Context) {
	var (
		user       models.UserAuth
		userDetail models.UserDetail
		err        error
	)
	if err = c.ShouldBindJSON(&user); err != nil {
		common.GenResponse(c, consts.BindingJsonErr, nil, "bing json failed")
		return
	}
	if userDetail, err = ctl.ServiceCheckUser(user.Username, user.Password); err != nil || userDetail.ID == 0 {
		common.GenResponse(c, consts.UserNameOrPasswordErr, nil, "username or password error")
		return
	}
	response := make(map[string]interface{})
	response["user_detail"] = userDetail
	response["token"] = middles.GenerateJWT(userDetail.Username,
		[]string{}, []string{}, userDetail.ID, userDetail.IsAdmin)
	common.GenResponse(c, consts.Success, response, "")

}
