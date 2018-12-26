// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package mocks

import (
	"{{.ProjectPath}}/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

type Mock{{.ModelName}}Repository struct {
	mock.Mock
}
 
func (m *Mock{{.ModelName}}Repository) Find{{.ModelName}}ByID(DB *gorm.DB, id int) (model models.{{.ModelName}}Detail, err error) {
	ret := m.Called(id)
	if rf, ok := ret.Get(0).(func(int) models.{{.ModelName}}Detail); ok {
		model = rf(id)
	} else {
		model = ret.Get(0).(models.{{.ModelName}}Detail)
	}
	return
}
func (m *Mock{{.ModelName}}Repository) Insert{{.ModelName}}(DB *gorm.DB, model models.{{.ModelName}}Create) (model models.{{.ModelName}}Detail, err error) {
	return
}

// 删除 
func (m *Mock{{.ModelName}}Repository) Delete{{.ModelName}}(DB *gorm.DB, ids []int) (err error) {
	return
}

// 查询 
func (m *Mock{{.ModelName}}Repository) FindAll{{.ModelName}}(DB *gorm.DB, page, size int, order string, query string, queryArgs ...interface{}) (models []*models.{{.ModelName}}List, total int, err error) {
	return
}
