// Copyright 2018 itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package repositories

import (
	"{{.ProjectPath}}/pkg/models"
	"github.com/jinzhu/gorm"
)

type I{{.ModelName}}Repository interface {
	//根据ID查找
	Find{{.ModelName}}ByID(DB *gorm.DB, id int) (role models.{{.ModelName}}Detail, err error)
	// 创建
	Insert{{.ModelName}}(DB *gorm.DB, role models.{{.ModelName}}Create) (result models.{{.ModelName}}Detail, err error)
	// 修改
	Update{{.ModelName}}(DB *gorm.DB, role models.{{.ModelName}}Update) (result models.{{.ModelName}}Detail, err error)
	// 删除
	Delete{{.ModelName}}(DB *gorm.DB, ids []int) error
	// 查询
	FindAll{{.ModelName}}(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (results []*models.{{.ModelName}}List, total int, err error)
}
