// Copyright 2018 itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package repositories

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type IFileUploadRepository interface {
	//根据ID查找
	FindFileUploadByID(DB *gorm.DB, id int) (role models.FileUploadDetail, err error)
	//根据Hash查找
	FindFileUploadByHash(DB *gorm.DB, hashStr string) (role models.FileUploadDetail, err error)
	// 创建
	InsertFileUpload(DB *gorm.DB, role models.FileUploadCreate) (result models.FileUploadDetail, err error)
	// 删除
	DeleteFileUpload(DB *gorm.DB, ids []int) error
	// 查询
	FindAllFileUpload(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (results []*models.FileUploadList, total int, err error)
}
