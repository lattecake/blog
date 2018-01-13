package helper

import (
	"github.com/astaxie/beego"
	"strings"
)

func ReplaceImageUrl(images []string) string {
	for _, image := range images {
		return strings.Replace(image, "/mnt/storage/", beego.AppConfig.String("source_url"), 1)
	}
	return string("")
}

func ReplaceImageSrc(url string) string {
	return strings.Replace(strings.Replace(url, "/mnt/storage/uploads/", "/mnt/storage/", 1), "/mnt/storage/", beego.AppConfig.String("source_url"), 1)
}

func PostImageSrc(images map[int64]string, postId int64) string {
	return strings.Replace(images[postId], "/mnt/storage/uploads/", beego.AppConfig.String("source_url"), 1)
}
