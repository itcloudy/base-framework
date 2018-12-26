// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type FileUploadRepository struct {
}

//根据ID查找
func (repo *FileUploadRepository) FindFileUploadByID(DB *gorm.DB, id int) (model models.FileUploadDetail, err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("FindFileUploadByID failed", zap.Error(err), zap.Int("id", id))
		}
	}()
	err = DB.Where("id = ?", id).First(&model).Error

	return
}

//根据Hash查找
func (repo *FileUploadRepository) FindFileUploadByHash(DB *gorm.DB, hashStr string) (model models.FileUploadDetail, err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("FindFileUploadByHash failed", zap.Error(err), zap.String("hashStr", hashStr))
		}
	}()
	err = DB.Where("address LIKE ?", "%"+hashStr+"%").Limit(1).First(&model).Error
	logs.Logger.Error("InsertFileUpload failed", zap.Error(err), zap.Any("hashStr", model))
	return
}

// 创建
func (repo *FileUploadRepository) InsertFileUpload(DB *gorm.DB, model models.FileUploadCreate) (result models.FileUploadDetail, err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("InsertFileUpload failed", zap.Error(err), zap.Any("hashStr", model))
		}
	}()
	err = DB.Create(&model).Error
	if err == nil {
		result, err = repo.FindFileUploadByID(DB, model.ID)
	}
	return
}

// 删除
func (repo *FileUploadRepository) DeleteFileUpload(DB *gorm.DB, ids []int) (err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("DeleteFileUpload failed", zap.Error(err), zap.Ints("ids", ids))
		}
	}()
	err = DB.Where("id IN (?)", ids).Delete(models.FileUploadDetail{}).Error
	return
}

// 查询
func (repo *FileUploadRepository) FindAllFileUpload(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (results []*models.FileUploadList, total int, err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("FindAllFileUpload failed", zap.Error(err), zap.Int("page", page), zap.Int("size", size),
				zap.String("order", order), zap.String("query", query), zap.Any("queryArgs", queryArgs))
		}
	}()
	if len(order) == 0 {
		order = "id desc"
	}
	db := DB.Model(&models.FileUploadList{}).Order(order)
	if len(query) > 0 {
		db.Where(query, queryArgs[:])
	}
	db.Count(&total)
	err = db.Offset((page - 1) * size).Limit(size).Find(&results).Error
	return
}
