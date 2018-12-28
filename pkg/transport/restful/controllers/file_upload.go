// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/interfaces/services"
	"github.com/itcloudy/base-framework/pkg/transport/restful/common"
	"github.com/itcloudy/base-framework/tools"
)

type FileUploadController struct {
	services.IFileUploadService
}

//文件上传
func (ctl FileUploadController) CtlCreateFileUpload(c *gin.Context) {
	if file, header, err := c.Request.FormFile("upload_file"); file != nil && header != nil {
		result, err := ctl.ServiceFileUploadCreate(file, header.Filename, c.GetInt(consts.LoginUserID))
		if err == nil {
			common.GenResponse(c, consts.Success, result, "")
		} else {
			common.GenResponse(c, consts.FilelUploadErr, "", err.Error())
		}
	} else {
		common.GenResponse(c, consts.GetFileErr, "", err.Error())
	}

}

//多个文件上传
func (ctl FileUploadController) CtlCreateMultiFileUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		common.GenResponse(c, consts.FilelUploadErr, "", err.Error())

	}
	files := form.File["upload_files[]"]
	if len(files) > 0 {
		ctl.ServiceFileUploadMultiCreate(files, c.GetInt(consts.LoginUserID))
	} else {
		common.GenResponse(c, consts.FilelUploadErr, "", "upload files is empty")

	}

}

//根据ID获得详情
func (ctl FileUploadController) CtlGetFileUploadByID(c *gin.Context) {
	result, err := ctl.ServiceGetFileUploadByID(tools.StrToInt(c.Param("id")))
	if err != nil {
		common.GenResponse(c, consts.GetFileErr, "", err.Error())
	}
	common.GenResponse(c, consts.Success, result, "")
}
func (ctl FileUploadController) CtlGetFileUploadByHash(c *gin.Context) {
	result, err := ctl.ServiceGetFileUploadByHash(c.Param("hash"))
	if err != nil {
		common.GenResponse(c, consts.GetFileErr, "", err.Error())
	}
	common.GenResponse(c, consts.Success, result, "")
}

//更新
func (ctl FileUploadController) CtlUpdateFileUploadByID(c *gin.Context) {
	result, err := ctl.ServiceGetFileUploadByID(tools.StrToInt(c.Param("id")))
	if err != nil {
		common.GenResponse(c, consts.GetFileErr, "", err.Error())
	}
	common.GenResponse(c, consts.Success, result, "")
}

//查询
func (ctl FileUploadController) CtlGetAllFileUpload(c *gin.Context) {
	var (
		page  int
		size  int
		order string
	)
	result := make(map[string]interface{})
	size = tools.StringToIntDefault(c.Query("pageSize"), consts.DefaultSize)
	page = tools.StringToIntDefault(c.Query("currentPage"), consts.DefaultPage)
	list, pagination, err := ctl.ServiceGetAllFileUpload(page, size, order, "")
	if err != nil {

	}
	result["pagination"] = pagination
	result["list"] = list
	common.GenResponse(c, consts.Success, result, "")
}
