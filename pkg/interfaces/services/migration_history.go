// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import "github.com/itcloudy/base-framework/pkg/models"

type IMigrationHistory interface {
	ServiceGetCurrentVersion() (version string, err error)
	ServiceFirstMigration() (err error)
	ServiceUpdateToOneVersion(version string) (err error)
	ServiceGetAllListMigration() (migrates []models.MigrationHistory, err error)
}
