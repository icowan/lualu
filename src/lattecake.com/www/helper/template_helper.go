package helper

import (
	"strings"
	"lattecake.com/www/models"
	"github.com/astaxie/beego"
)

func ReplaceImageUrl(images []*models.Image) string {
	for _, image := range images {
		return strings.Replace(image.RealPath, "/mnt/storage/", beego.AppConfig.String("source_url"), 1)
	}
	return string("")
}

func ReplaceImageSrc(url string) string {
	return strings.Replace(url, "/mnt/storage/uploads/", beego.AppConfig.String("source_url"), 1)
}