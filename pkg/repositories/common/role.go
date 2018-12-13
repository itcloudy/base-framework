// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
	"strconv"
)

type RoleRepository struct {
}

//根据ID查找
func (repo *RoleRepository) FindRoleByID(DB *gorm.DB, id string) (role models.RoleDetail, err error) {
	err = DB.Where("id = ?", id).First(&role).Error
	return
}

// 创建角色
func (repo *RoleRepository) InsertRole(DB *gorm.DB, model models.RoleCreate) (result models.RoleDetail, err error) {
	model.ID = 0
	err = DB.Create(&model).Error
	if err == nil {
		return repo.FindRoleByID(DB, strconv.Itoa(model.ID))
	}
	return
}

// 修改角色
func (repo *RoleRepository) UpdateRole(DB *gorm.DB, role models.RoleUpdate) (result models.RoleDetail, err error) {
	err = DB.Updates(role).Error
	return
}

// 删除角色
func (repo *RoleRepository) DeleteRole(DB *gorm.DB, ids []string) (err error) {
	return DB.Where("id IN (?)", ids).Delete(models.RoleDetail{}).Error
}

// 查询角色
func (repo *RoleRepository) FindAllRole(DB *gorm.DB, offset, limit int, order string, query string, queryArgs ...interface{}) (roles []*models.RoleList, count int, err error) {
	err = DB.Order(order).Offset(offset).Limit(limit).Find(&roles).Error
	return
}
