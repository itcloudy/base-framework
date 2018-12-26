// Copyright 2018 itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/interfaces/repositories"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/tools"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FilesUploadService struct {
	DB *gorm.DB
	repositories.IFileUploadRepository
}

func (service *FilesUploadService) ServiceGetFileUploadByID(id int) (result models.FileUploadDetail, err error) {
	return service.FindFileUploadByID(service.DB, id)

}
func (service *FilesUploadService) ServiceGetFileUploadByHash(hashStr string) (role models.FileUploadDetail, err error) {
	return service.FindFileUploadByHash(service.DB, hashStr)

}
func (service *FilesUploadService) ServiceFileUploadDelete(ids []int) (err error) {
	return service.DeleteFileUpload(service.DB, ids)
}
func (service *FilesUploadService) ServiceGetAllFileUpload(page, size int, order string, query string, queryArgs ...interface{}) (results []*models.FileUploadList, pagination conf.Pagination, err error) {
	pagination.Current = page
	pagination.Size = size
	results, pagination.Total, err = service.FindAllFileUpload(service.DB, page, size, order, query, queryArgs)
	return
}

func (service *FilesUploadService) ServiceFileUploadCreate(file multipart.File, fileName string, ownId int) (result models.FileUploadDetail, err error) {
	var fileTarget string
	target := conf.Config.FileUpload.Target
	switch target {
	case "local":
		fileTarget, err = localStorage(file, fileName)
		break
	case "qiniu":
		fileTarget, err = qiNiuStorage(file, fileName)
		break
	default:
		fileTarget, err = localStorage(file, fileName)
	}
	if err == nil {
		model := models.FileUploadCreate{UserID: ownId, FileName: fileName, Type: target}
		if target == "local" {
			model.Address = path.Join(consts.USER_UPLOAD_FILE_URL, fileTarget)
		}
		//判断是否存在
		if result, err = service.FindFileUploadByHash(service.DB, model.Address); err == nil && result.ID > 0 {
			return
		}
		result, err = service.InsertFileUpload(service.DB, model)
		if err != nil {
			logs.Logger.Error("file upload success but record insert db occur err", zap.Error(err), zap.String("file target", target))
		}

	} else {
		logs.Logger.Error("file upload failed", zap.Error(err), zap.String("file target", target))
	}
	return
}
func (service *FilesUploadService) ServiceFileUploadMultiCreate(files []*multipart.FileHeader, ownId int) (result map[string]string, err error) {
	//判断文件夹是否存在
	uploadPath := conf.Config.UserDataDir
	if exists, _ := tools.PathExists(uploadPath); !exists {
		err := os.Mkdir(uploadPath, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	for _, uploadFile := range files {
		filename := uploadFile.Filename
		src, err := uploadFile.Open()
		defer src.Close()
		if err != nil {
			result[filename] = err.Error()
			continue
		}
		var bys []byte
		databuf := bytes.NewBuffer(bys)
		for {
			var oneBuff = make([]byte, 1024)
			oneLen, _ := src.Read(oneBuff)
			if oneLen > 0 {
				databuf.Write(oneBuff)
			} else {
				break
			}
		}
		src.Seek(0, 0)
		hash := md5.New()
		hash.Write(databuf.Bytes())
		hashName := hex.EncodeToString(hash.Sum(nil))
		sufixList := strings.Split(filename, ".")
		sufix := sufixList[len(sufixList)-1]
		newName := tools.StringsJoin(hashName, ".", sufix)
		filePath := path.Join(uploadPath, newName)
		out, err := os.Create(filePath)

		defer out.Close()
		if _, err = io.Copy(out, src); err != nil {
			result[filename] = err.Error()
		} else {
			result[filename] = newName
		}
	}
	return
}
func localStorage(file multipart.File, fileName string) (result string, err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("localStorage", zap.Error(err))
		}
	}()
	var (
		bys    []byte
		exists bool
	)
	//判断文件夹是否存在
	uploadPath := conf.Config.UserDataDir
	exists, _ = tools.PathExists(uploadPath)

	if !exists {
		err := os.Mkdir(uploadPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	databuf := bytes.NewBuffer(bys)
	for {
		var oneBuff = make([]byte, 1024)
		oneLen, _ := file.Read(oneBuff)
		if oneLen > 0 {
			databuf.Write(oneBuff)
		} else {
			break
		}
	}
	file.Seek(0, 0)
	hash := md5.New()
	hash.Write(databuf.Bytes())
	hashName := hex.EncodeToString(hash.Sum(nil))
	sufixList := strings.Split(fileName, ".")
	sufix := sufixList[len(sufixList)-1]
	result = tools.StringsJoin(hashName, ".", sufix)
	filePath := path.Join(uploadPath, result)

	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return
}
func qiNiuStorage(file multipart.File, fileName string) (result string, err error) {
	defer func() {
		if err != nil {
			logs.Logger.Error("qiNiuStorage", zap.Error(err))
		}
	}()
	return
}
