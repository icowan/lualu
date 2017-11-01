package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"lattecake.com/fig-bed/core"
	"crypto/sha1"
	"io"
	"log"
	"encoding/hex"
	"mime/multipart"
	"time"
	"os"
)

var config *core.Config

func init() {

	config = core.Instance(".env")

	switch config.Env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	}
}

func main() {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Any("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://lattecake.com")
	})

	router.POST("/putImage", UploadImage)

	router.Run(":8089")
}

func UploadImage(c *gin.Context) {
	//token := c.Request.Form.Get("token")
	token := c.Request.FormValue("token")
	if token == "" || token != config.UploadImageToken {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": "error",
			"msg":  "token不正确",
		})
		fmt.Println(map[string]interface{}{
			"code":                    "error",
			"msg":                     "token不正确",
			"token":                   token,
			"config.UploadImageToken": config.UploadImageToken,
		})
		return
	}


	file, header, err := c.Request.FormFile("MarkdownImage")
	filename := header.Filename


	fmt.Println(header.Filename)
	out, err := os.Create("./tmp/"+filename+".png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}




	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": "error",
			"msg":  fmt.Sprintf("get form err: %s", err.Error()),
		})
		fmt.Println(map[string]interface{}{
			"code": "error",
			"msg":  fmt.Sprintf("get form err: %s", err.Error()),
		})
		return
	}

	files := form.File["MarkdownImage"]

	var fileInfo *multipart.FileHeader
	var fileName string

	path := time.Now().Format("2006/01/02")

	//filePath := config.UploadImageDir + path

	for _, file := range files {

		h := sha1.New()
		f, _ := file.Open()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		f.Close()

		fileName = hex.EncodeToString(h.Sum(nil))

		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code": "error",
				"msg":  fmt.Sprintf("upload file err: %s", err.Error()),
			})
			fmt.Println(map[string]interface{}{
				"code": "error",
				"msg":  fmt.Sprintf("upload file err: %s", err.Error()),
			})
			continue
		}
		fileInfo = file
		break
	}

	//c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields token=%s", len(files), token))

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": "success",
		"data": map[string]interface{}{
			"width":     0,
			"height":    0,
			"filename":  fileInfo.Filename,
			"storename": fileName,
			"size":      0,
			"path":      "",
			"hash":      "",
			"timestamp": time.Now().Unix(),
			"url":       config.UploadImageUrl + "uploads/images/" + path + "/" + fileInfo.Filename,
			"delete":    "",
		},
	})
}
