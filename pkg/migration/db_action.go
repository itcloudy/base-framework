// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package migration

import (
	"github.com/itcloudy/base-framework/pkg/migration/mysql"
	"github.com/itcloudy/base-framework/pkg/migration/postgres"
	"github.com/itcloudy/base-framework/pkg/models"
)

var AllInitMigrations = map[string][]models.MigrationHistory{
	"postgres": {
		models.MigrationHistory{
			Version: "0.0.1",
			Data:    postgres.Init,
		},
	},
	"mysql": {
		models.MigrationHistory{
			Version: "0.0.1",
			Data:    mysql.Init,
		},
	},
}
var AllUpdateMigrations = map[string][]models.MigrationHistory{
	"postgres": {
		models.MigrationHistory{
			Version: "0.0.2",
			Data:    postgres.Update_0_0_2,
		},
	},
	"mysql": {},
}
