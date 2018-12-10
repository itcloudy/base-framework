// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
	"strconv"
)

type SystemAPIRepository struct {
}

//根据ID查找
func (repo *SystemAPIRepository) FindSystemAPIByID(DB *gorm.DB, id string) (api models.SystemApiDetail, err error){
	err = DB.Where("id = ?", id).First(&api).Error
	return
}
// 创建系统接口
func (repo *SystemAPIRepository) InsertSystemAPI(DB *gorm.DB, create models.SystemApiCreate) (result models.SystemApiDetail, err error){
	err = DB.Create(&create).Error
	if err == nil {
		return repo.FindSystemAPIByID(DB, strconv.Itoa(create.ID))
	}
	return
}
// 修改系统接口
func (repo *SystemAPIRepository) UpdateSystemAPI(DB *gorm.DB, api models.SystemApiUpdate) (result models.SystemApiDetail, err error){
	err = DB.Updates(api).Error
	return
}
// 获得所有系统接口
func (repo *SystemAPIRepository) FindAllSystemAPI(DB *gorm.DB) (apis []*models.SystemApiDetail, err error){
	err = DB.Find(&apis).Error
	return
}