// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package repositories

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type ISystemAPIRepository interface {
	//根据ID查找
	FindSystemAPIByID(DB *gorm.DB, id string) (api models.SystemApiDetail, err error)
	// 创建系统接口
	InsertSystemAPI(DB *gorm.DB, api models.SystemApiCreate) (result models.SystemApiDetail, err error)
	// 修改系统接口
	UpdateSystemAPI(DB *gorm.DB, api models.SystemApiUpdate) (result models.SystemApiDetail, err error)
	// 删除系统接口
	DeleteSystemAPI(DB *gorm.DB, ids []string) error
	// 查询系统接口
	FindAllSystemAPI(DB *gorm.DB, offset, limit int, order string, query string, queryArgs ...interface{}) (apis []*models.SystemApiList, count int, err error)
}
