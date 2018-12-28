// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/models"
	"mime/multipart"
)

type IFileUploadService interface {
	ServiceGetFileUploadByID(id int) (result models.FileUploadDetail, err error)
	ServiceGetFileUploadByHash(hashStr string) (result models.FileUploadDetail, err error)
	ServiceFileUploadCreate(file multipart.File, fileName string, ownId int) (result models.FileUploadDetail, err error)
	ServiceFileUploadMultiCreate(files []*multipart.FileHeader, ownId int) (result map[string]string, err error)
	ServiceFileUploadDelete(ids []int) (err error)
	ServiceGetAllFileUpload(page, size int, order string, query string, queryArgs ...interface{}) (results []*models.FileUploadList, pagination conf.Pagination, err error)
}
