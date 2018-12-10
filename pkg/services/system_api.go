// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/interfaces/repositories"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type SystemAPIService struct {
	DB *gorm.DB
	repositories.ISystemAPIRepository
}

func (service *SystemAPIService) ServiceGetAllSystemAPI() (systemApis []*models.SystemApiList, err error) {
	err = service.DB.Find(&systemApis).Error
	return
}
