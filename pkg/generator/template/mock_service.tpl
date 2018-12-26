// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"{{.ProjectPath}}/pkg/conf"
	"{{.ProjectPath}}/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

type Mock{{.ModelName}}Service struct {
	mock.Mock
}

//根据ID查找
func (mo *Mock{{.ModelName}}Service) Get{{.ModelName}}ByID(id int) (model models.{{.ModelName}}Detail, err error) {
	ret := mo.Called(id)
	if rf, ok := ret.Get(0).(func(int) models.{{.ModelName}}Detail); ok {
		model = rf(id)
	} else {
		model = ret.Get(0).(models.{{.ModelName}}Detail)
	}
	return model, err
}

// 创建
func (mo *Mock{{.ModelName}}Service) Insert{{.ModelName}}(DB *gorm.DB, model models.{{.ModelName}}Create) (result models.{{.ModelName}}Detail, err error) {
	return
}

// 删除
func (mo *Mock{{.ModelName}}Service) Delete{{.ModelName}}(DB *gorm.DB, ids []int) (err error) {
	return
}

// 查询系统接口
func (mo *Mock{{.ModelName}}Service) FindAll{{.ModelName}}(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (models []*models.{{.ModelName}}List, pagination conf.Pagination, err error) {
	return
}
