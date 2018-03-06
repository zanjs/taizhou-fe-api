package handler

import (
	"io"
	"os"
	"path"
	"strings"

	"anla.io/taizhou-fe-api/config"
	"anla.io/taizhou-fe-api/response"
	"github.com/houndgo/houndgo/ifile"
	"github.com/houndgo/houndgo/itime"
	"github.com/houndgo/suuid"
	"github.com/kataras/iris"
)

// UploadFile is ...
func UploadFile(ctx iris.Context) {
	file, info, err := ctx.FormFile("file")

	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	defer file.Close()
	fname := info.Filename
	fileExt := path.Ext(fname)
	confUp := config.Config.Upload
	conffileEx := confUp.Ext
	ofbool := strings.Contains(conffileEx, fileExt)
	if ofbool != true {
		response.JSONError(ctx, "不支持该格式")
		return
	}
	today := itime.Today()
	todayPath := confUp.Path + "/" + today
	checkBool := ifile.CheckFileIsExist(todayPath)
	if !checkBool {
		ifile.Mkdir(todayPath)
	}

	sfileName := suuid.New().String() + fileExt
	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	newFilePath := todayPath + "/" + sfileName
	out, err := os.OpenFile(newFilePath,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		response.JSONError(ctx, err.Error())
		return
	}
	defer out.Close()

	io.Copy(out, file)

	response.JSON(ctx, newFilePath)
}
