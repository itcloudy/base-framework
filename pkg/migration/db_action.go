// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package migration

import (
	"github.com/itcloudy/base-framework/pkg/migration/postgres"
	"github.com/itcloudy/base-framework/pkg/models"
)

var AllInitMigrations = []models.MigrationHistory{
	models.MigrationHistory{
		Version: "0.0.1",
		Data:    postgres.Init,
	},
}
var AllUpdateMigrations []models.MigrationHistory
