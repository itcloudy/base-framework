// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package mocks

import (
	"github.com/hashicorp/go-version"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/stretchr/testify/mock"
)

type IMigrationHistoryRepository struct {
	mock.Mock
}

func (_m *IMigrationHistoryRepository) CurrentVersion() (version string, err error) {
	return
}
func (_m *IMigrationHistoryRepository) ApplyMigrations(collection version.Collection, migrates map[string]string) (err error) {
	return
}
func (_m *IMigrationHistoryRepository) ListMigration() (migrates []models.MigrationHistory, err error) {
	return
}
