// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/interfaces/repositories"
	"github.com/itcloudy/base-framework/pkg/migration"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/pkg/repositories/common"
	"github.com/itcloudy/base-framework/tools"
	"github.com/jmoiron/sqlx"
)

// 数据库升级服务
type MigrationService struct {
	DB *sqlx.DB
	repositories.IMigrationHistoryRepository
}

/*
获得已升级的最新版本
*/
func (service *MigrationService) ServiceGetCurrentVersion() (version string, err error) {
	return service.CurrentVersion(service.DB)

}

/*
第一次初始化
*/
func (service *MigrationService) ServiceFirstMigration() (err error) {
	// 删除数据库中的表
	err = dropTables()
	if err != nil {
		return
	}
	var (
		collection version.Collection
	)
	needUpdateMap := make(map[string]string)

	// 判断需要升级几个版本
	for _, migrate := range migration.AllInitMigrations[conf.Config.DB.DbType] {
		var migVersion *version.Version
		// 如果转换失败则返回
		if migVersion, err = version.NewVersion(migrate.Version); err != nil {
			return
		}
		collection = append(collection, migVersion)
		needUpdateMap[migrate.Version] = migrate.Data
	}
	// 调用repository
	if len(collection) < 1 {
		return errors.New("no need update version, code logic has problem")
	}
	err = service.ApplyMigrations(service.DB, collection, needUpdateMap)
	if err == nil {
		// 插入超级管理员
		var superUser models.UserCreate
		cfg := conf.Config.DB
		conf.GetDBConnection(cfg.DbType, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.Charset, "")
		userService := UserService{}
		switch cfg.DbType {
		case "mysql":
		case "postgres":
			userService.IUserRepository = &common.UserRepository{}
			break
		default:
			panic(errors.New("un support sql type:" + cfg.DbType))

		}
		userService.DB = conf.DBConn
		admin := conf.Config.Admin
		superUser.Username = admin.UserName
		superUser.Password = admin.Password
		superUser.ConfirmPassword = admin.Password
		superUser.Email = admin.Email
		superUser.Mobile = admin.Mobile
		superUser.IsActive = true
		superUser.IsAdmin = true
		_, err = userService.ServiceUserCreate(superUser)
	}
	conf.DBConn.Close()
	return
}

/*
升级到某个版本，若中间存在多个，则中间版本同样升级
*/
func (service *MigrationService) ServiceUpdateToOneVersion(ver string) (err error) {
	var (
		collection version.Collection

		needVer     *version.Version
		last        *version.Version
		lastVersion string
	)
	needUpdateMap := make(map[string]string)
	if needVer, err = version.NewVersion(ver); err != nil {
		return
	}

	// 获得已经安装的最新版本
	if lastVersion, err = service.CurrentVersion(service.DB); err != nil {
		return
	}
	if last, err = version.NewVersion(lastVersion); err != nil {
		return
	}

	// 判断升级的版本和已升级的最后一个的大小
	if needVer.LessThan(last) || needVer.Equal(last) {
		return errors.New("need update version:  " + conf.Config.DBUpdateToVersion + " last installed version: : " + last.String())
	}
	// 判断需要升级几个版本
	for _, migrate := range migration.AllUpdateMigrations[conf.Config.DB.DbType] {
		var migVersion *version.Version
		// 如果转换失败则返回
		if migVersion, err = version.NewVersion(migrate.Version); err != nil {
			return
		}
		// 判断版本大小，只有大于最新版本，小于等于更新版本的才放进collection中
		if (migVersion.LessThan(needVer) && migVersion.GreaterThan(last)) || migVersion.Equal(needVer) && migVersion.GreaterThan(last) {
			collection = append(collection, migVersion)
			needUpdateMap[migrate.Version] = migrate.Data
		}
	}
	// 调用repository
	if len(collection) < 1 {
		return errors.New("no need update version, code logic has problem")
	}
	return service.ApplyMigrations(service.DB, collection, needUpdateMap)
}

/*
列出所有的版本，包括系统中存在的没有安装的
*/
func (service *MigrationService) ServiceGetAllListMigration() (results []models.MigrationHistory, err error) {
	var installedMigrates []models.MigrationHistory
	results = migration.AllInitMigrations[conf.Config.DB.DbType]
	results = append(results, migration.AllUpdateMigrations[conf.Config.DB.DbType]...)
	var (
		verSlice []string
	)
	// 获得已经安装的版本
	installedMigrates, _, _ = service.ListMigration(service.DB)
	//获得已安装的版本
	for _, migrate := range installedMigrates {
		verSlice = append(verSlice, migrate.Version)
	}
	//判断已经安装的版本
	for k, migrate := range results {
		if tools.StringInSlice(verSlice, migrate.Version) {
			results[k].Installed = true
		}
	}
	return
}

// dropTables is dropping all of the tables
func dropTables() (err error) {
	dbType := conf.Config.DB.DbType
	db := conf.SqlxDB
	if dbType == "postgres" {

		_, err = db.Exec(`
		DO $$ DECLARE
	    	r RECORD;
		BEGIN
	    	FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
			EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
	    	END LOOP;
		END $$;
		`)
		return err
	} else if dbType == "mysql" {
		// delete foreign key
		db.Exec(`DROP  PROCEDURE IF  EXISTS procedure_drop_foreign_key;`)
		_, er := db.Exec(fmt.Sprintf(`
CREATE PROCEDURE procedure_drop_foreign_key()
BEGIN
  DECLARE DB_NAME varchar(50) DEFAULT "%s"; 
  DECLARE done INT DEFAULT 0;
  DECLARE tableName varchar(50);   
  DECLARE constraintName varchar(50);   
  DECLARE cmd varchar(450);         
  DECLARE sur CURSOR               
  FOR 
  
  SELECT   TABLE_NAME , CONSTRAINT_NAME 
  FROM information_schema.key_column_usage 
  WHERE CONSTRAINT_SCHEMA = DB_NAME 
  AND referenced_table_name IS NOT NULL;
  DECLARE CONTINUE HANDLER FOR SQLSTATE '02000' SET done = 1;
 
  OPEN sur;
  REPEAT
    FETCH sur INTO tableName,constraintName;
    IF NOT done THEN 
		set cmd=concat('ALTER TABLE ', tableName, ' DROP FOREIGN KEY ', constraintName);
        SET @E=cmd; 
        PREPARE stmt FROM @E; 
          EXECUTE stmt;  
        DEALLOCATE PREPARE stmt;  
    END IF;
  UNTIL done END REPEAT;
  CLOSE sur;
END;`, conf.Config.DB.Name))
		if er != nil {
			return er
		}
		_, err = db.Exec(`call procedure_drop_foreign_key();`)
		if err != nil {
			return err
		}
		// drop all tables
		db.Exec(`DROP  PROCEDURE IF  EXISTS procedure_drop_table;`)
		db.Exec(fmt.Sprintf(`
CREATE PROCEDURE procedure_drop_table()
BEGIN
  DECLARE DB_NAME varchar(50) DEFAULT "%s"; 
  DECLARE done INT DEFAULT 0;
  DECLARE tableName varchar(50);   
  DECLARE cmd varchar(50);         
  DECLARE sur CURSOR               
  FOR 
  SELECT table_name FROM information_schema.TABLES WHERE table_schema=DB_NAME; 
  DECLARE CONTINUE HANDLER FOR SQLSTATE '02000' SET done = 1;
 
  OPEN sur;
  REPEAT
    FETCH sur INTO tableName;
    IF NOT done THEN 
       set cmd=concat('DROP TABLE ',DB_NAME,'.',tableName);    
        SET @E=cmd; 
        PREPARE stmt FROM @E; 
          EXECUTE stmt;  
         DEALLOCATE PREPARE stmt;  
    END IF;
  UNTIL done END REPEAT;
  CLOSE sur;
END;

`, conf.Config.DB.Name))
		_, err = db.Exec(`call procedure_drop_table();`)
		return err
	} else {
		return errors.New("db type not support")
	}
}
