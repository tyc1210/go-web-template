package v1

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"go-web-template/global"
	"go-web-template/pkg/app"
	"go-web-template/pkg/errcode"
	"go-web-template/pkg/logger"
	"go-web-template/pkg/util"
	"strings"
)

type File struct {
}

func NewFile() File {
	return File{}
}

// Upload2Local
// @Summary 上传
// @Tags 文件管理
// @Accept multipart/form-data
// @Produce  json
// @Param files formData file true "file"
// @Success 200 {object} app.Response "成功"
// @Router /api/upload [post]
func (f File) Upload2Local(c *gin.Context) {
	result := app.NewCommonResult(c)
	form, err := c.MultipartForm()
	if err != nil {
		logger.Error("Error when try to get file: %v", err)
		result.Error(&errcode.FileError)
	}
	// 获取所有文件
	files := form.File["files"]
	data := make([]string, 0, len(files))
	for _, file := range files {
		fileName := strings.Split(file.Filename, ".")
		name := fileName[0]
		ext := fileName[1]
		allExt := mapset.NewSetFromSlice(global.Cfg.App.UploadAllowExt)
		if !allExt.Contains(ext) {
			result.Error(&errcode.FileErrorExt)
			return
		}
		name = util.EncodeMD5(name) + "." + ext
		c.SaveUploadedFile(file, global.Cfg.App.UploadSavePath+"/"+name)
		data = append(data, global.Cfg.App.UploadServerUrl+name)
	}
	result.Success(data)
}
