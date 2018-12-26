// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package routers

import (
	"github.com/itcloudy/base-framework/pkg/transport/restful/controllers"
	"sync"
)

type kernel struct{}

var (
	k             *kernel
	containerOnce sync.Once
)

func restContainer() iRestContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

type iRestContainer interface {
	IndexContainer() controllers.IndexController
	UserContainer() controllers.UserController
	RoleContainer() controllers.RoleController
	SystemAPIContainer() controllers.SystemAPIController
}
