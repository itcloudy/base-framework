// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/tools"
	"github.com/jinzhu/gorm"
)

type SystemAPIRepository struct {
}

//根据ID查找
func (repo *SystemAPIRepository) FindSystemAPIByID(DB *gorm.DB, id int) (api models.SystemApiDetail, err error) {
	err = DB.Where("id = ?", id).First(&api).Error
	return
}

// 创建系统接口
func (repo *SystemAPIRepository) InsertSystemAPI(DB *gorm.DB, model models.SystemApiCreate) (result models.SystemApiDetail, err error) {
	//合成display字段信息
	model.Display = tools.StringsJoin("[", model.Method, "] ", model.Name, " (", model.Address, ")")
	model.ID = 0
	err = DB.Create(&model).Error
	if err == nil {
		return repo.FindSystemAPIByID(DB, model.ID)
	}
	return
}

// 修改系统接口
func (repo *SystemAPIRepository) UpdateSystemAPI(DB *gorm.DB, api models.SystemApiUpdate) (result models.SystemApiDetail, err error) {
	//合成display字段信息
	api.Display = tools.StringsJoin("[", api.Method, "] ", api.Name, " (", api.Address, ")")
	err = DB.Updates(api).Error
	return
}

// 接口禁用可用
func (repo *SystemAPIRepository) ActiveSystemAPI(DB *gorm.DB, ids []int, active bool) (err error) {
	return DB.Model(&models.SystemApiDetail{}).Where("id IN (?)", ids).Update("IsActive", active).Error
}

// 删除系统接口
func (repo *SystemAPIRepository) DeleteSystemAPI(DB *gorm.DB, ids []int) (err error) {
	return DB.Where("id IN (?)", ids).Delete(models.SystemApiDetail{}).Error

}

// 查询系统接口
func (repo *SystemAPIRepository) FindAllSystemAPI(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (results []*models.SystemApiList, total int, err error) {
	if len(order) == 0 {
		order = "id desc"
	}
	db := DB.Model(&models.SystemApiList{}).Order(order)
	if len(query) > 0 {
		db = db.Where(query, queryArgs[:])
	}
	db.Count(&total)
	err = db.Offset((page - 1) * size).Limit(size).Find(&results).Error
	return
}
