// Copyright 2018  itcloudy@qq.com  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package routers

import (
	"{{.ProjectPath}}/pkg/conf"
	"{{.ProjectPath}}/pkg/repositories/common"
	"{{.ProjectPath}}/pkg/restful/controllers"
	"{{.ProjectPath}}/pkg/services"
)

func (k *kernel) IndexContainer() controllers.IndexController {
	return controllers.IndexController{}
}
{{.ContainerInstance}}
