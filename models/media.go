package models

import (
	"os"
	"strings"
	"unicode/utf8"

	"anla.io/taizhou-fe-api/config"
	"anla.io/taizhou-fe-api/utils"
	"github.com/houndgo/suuid"
)

// Image is
type Image struct {
	IDModel
	TitleModel
	OrignalTitle string `json:"orignalTitle"`
	URL          string `json:"url"`
	Width        uint   `json:"width"`
	Height       uint   `json:"height"`
	Mime         string `json:"mime"`
}

// ImageUploadedInfo is 图片上传后的相关信息(目录、文件路径、文件名、UUIDName、请求URL)
type ImageUploadedInfo struct {
	UploadDir      string
	UploadFilePath string
	Filename       string
	UUIDName       string
	ImgURL         string
}

// GenerateImgUploadedInfo 创建一个ImageUploadedInfo
func GenerateImgUploadedInfo(ext string) ImageUploadedInfo {
	sep := string(os.PathSeparator)
	uploadImgDir := ""
	length := utf8.RuneCountInString(uploadImgDir)
	lastChar := uploadImgDir[length-1:]
	ymStr := utils.GetTodayYM(sep)

	var uploadDir string
	if lastChar != sep {
		uploadDir = uploadImgDir + sep + ymStr
	} else {
		uploadDir = uploadImgDir + ymStr
	}

	uuidName := suuid.New().String()
	filename := uuidName + ext
	uploadFilePath := uploadDir + sep + filename
	imgURL := strings.Join([]string{
		config.Config.Media.Host + config.Config.Media.Path,
		ymStr,
		filename,
	}, "/")
	return ImageUploadedInfo{
		ImgURL:         imgURL,
		UUIDName:       uuidName,
		Filename:       filename,
		UploadDir:      uploadDir,
		UploadFilePath: uploadFilePath,
	}
}
