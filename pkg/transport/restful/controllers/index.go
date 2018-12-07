// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

// @tags  首页
// @Description 框架介绍
// @Summary 框架介绍
// @Accept  json
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"token":"Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjc1MzIxNjMwODMsIk5hbWUiOiJhZG1pbiIsIlJvbGUiOm51bGwsIlVzZXJJZCI6MiwiSXNBZG1pbiI6ZmFsc2V9.HZq5jBw4-ZQipQPnq0K7Ei0_LvaRXZGNgKqLoFnhV_vpfQupmddsDMZbiI_Yy0Zhd7J7AvRGDXMfVwW9-TidsDrux6-L4KQWIV0Mrlj4SXgW13HvMSXW0XzHYQBxiai61AeJx4VmQR84s2lI5hmKuiVOpsyOZAduJoO1K26b8X4","user":{"id":2,"name":"admin","alias":"","email":"","password":"","roles":[],"openid":"admin","active":true,"is_admin":false}},"message":"success"}"
// @Router /user/{id} [get]
func (ctl *IndexController) Information(c *gin.Context) {
	c.JSON(http.StatusOK, "golang framework")
}
