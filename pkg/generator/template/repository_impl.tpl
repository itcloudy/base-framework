// Copyright 2018 itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"{{.ProjectPath}}/pkg/logs"
	"{{.ProjectPath}}/pkg/models"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type {{.ModelName}}Repository struct {
}

//根据ID查找
func (repo *{{.ModelName}}Repository) Find{{.ModelName}}ByID(DB *gorm.DB, id int) (api models.{{.ModelName}}Detail, err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("Find{{.ModelName}}ByID failed", zap.Error(err), zap.Int("id", id))
		}
	}()
	err = DB.Where("id = ?", id).First(&api).Error
	return
}

// 创建
func (repo *{{.ModelName}}Repository) Insert{{.ModelName}}(DB *gorm.DB, model models.{{.ModelName}}Create) (result models.{{.ModelName}}Detail, err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("Insert{{.ModelName}} failed", zap.Error(err), zap.Any("model", model))
		}
	}()
	model.ID = 0
	err = DB.Create(&model).Error
	if err == nil {
		return repo.Find{{.ModelName}}ByID(DB, model.ID)
	}
	return
}

// 修改
func (repo *{{.ModelName}}Repository) Update{{.ModelName}}(DB *gorm.DB, model models.{{.ModelName}}Update) (result models.{{.ModelName}}Detail, err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("Update{{.ModelName}} failed", zap.Error(err), zap.Any("model", model))
		}
	}()
	err = DB.Updates(model).Error
	return
}

// 删除
func (repo *{{.ModelName}}Repository) Delete{{.ModelName}}(DB *gorm.DB, ids []int) (err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("Delete{{.ModelName}} failed", zap.Error(err), zap.Ints("ids", ids))
		}
	}()
	err = DB.Where("id IN (?)", ids).Delete(models.{{.ModelName}}Detail{}).Error
	return

}

// 查询
func (repo *{{.ModelName}}Repository) FindAll{{.ModelName}}(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (results []*models.{{.ModelName}}List, total int, err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("FindAll{{.ModelName}} failed", zap.Error(err), zap.Int("page", page), zap.Int("size", size),
				zap.String("order", order), zap.String("query", query), zap.Any("queryArgs", queryArgs))
		}
	}()
	if len(order) == 0 {
		order = "id desc"
	}
	db := DB.Model(&models.{{.ModelName}}List{}).Order(order)
	if len(query) > 0 {
		db.Where(query, queryArgs[:])
	}
	db.Count(&total)
	err = db.Offset((page - 1) * size).Limit(size).Find(&results).Error
	return
}
