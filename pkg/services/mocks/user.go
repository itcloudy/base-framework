// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package mocks

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/stretchr/testify/mock"
)

type IUserService struct {
	mock.Mock
}

func (_m *IUserService) GetSelf(id string) (user models.User, err error) {
	return
}
