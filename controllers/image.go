package controllers

import (
	"github.com/lattecake/blog/models"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"github.com/lattecake/utils"
	"time"
)

//  ImageController operations for Image
type ImageController struct {
	BaseController
}

// URLMapping ...
func (c *ImageController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Delete", c.Delete)
}

type resultImage struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
	Data struct {
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Filename  string `json:"filename"`
		Storename string `json:"storename"`
		Size      int64  `json:"size"`
		Path      string `json:"path"`
		Hash      string `json:"hash"`
		Timestamp int64  `json:"timestamp"`
		Url       string `json:"url"`
	} `json:"data"`
}

// Post ...
// @Title Post
// @Description create Image
// @Param	body		body 	models.Image	true		"body for Image content"
// @Success 201 {int} models.Image
// @Failure 403 body is empty
// @router /putImage [post]
func (c *ImageController) Post() {

	token := c.GetString("token")
	if len(token) < 1 {
		logs.Warn(c.WorkerId(), " | token is not exists.")
		c.Data["json"] = &resultImage{
			Code:    "fail",
			Message: "token is not exists.",
		}
		c.ServeJSON()
		return
	}

	if token != beego.AppConfig.String("upload_image_token") {
		logs.Warn(c.WorkerId(), " | token is error.")
		c.Data["json"] = &resultImage{
			Code:    "fail",
			Message: "token is error.",
		}
		c.ServeJSON()
		return
	}

	f, h, err := c.GetFile("MarkdownImage")
	if err != nil {
		logs.Error(c.WorkerId(), " | ", err)
		c.Data["json"] = &resultImage{
			Code:    "fail",
			Message: "file is not exists.",
		}
		c.ServeJSON()
		return
	}

	defer f.Close()

	md5FileName, err := utils.Md5File(f)
	if err != nil {
		logs.Error(c.WorkerId(), " | ", err)
		c.Data["json"] = &resultImage{
			Code:    "fail",
			Message: err.Error(),
		}
		c.ServeJSON()
		return
	}

	if img, err := models.GetImageByMd5(md5FileName); err == nil {
		// 已经存在 返回图片地址
		resImg := &resultImage{
			Code: "success",
			Data: struct {
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				Filename  string `json:"filename"`
				Storename string `json:"storename"`
				Size      int64  `json:"size"`
				Path      string `json:"path"`
				Hash      string `json:"hash"`
				Timestamp int64  `json:"timestamp"`
				Url       string `json:"url"`
			}{Width: 0, Height: 0, Filename: h.Filename, Storename: h.Filename, Size: h.Size, Path: "/", Hash: md5FileName, Timestamp: time.Now().Unix(), Url: beego.AppConfig.String("upload_image_url") + "/images/" + img.ImagePath},
		}

		c.Data["json"] = resImg
		c.ServeJSON()
		return
	}

	filePath := time.Now().Format("2006/01/") + md5FileName[len(md5FileName)-5:len(md5FileName)-3] + "/" + md5FileName[24:26] + "/" + md5FileName[16:17] + md5FileName[12:13] + "/"

	path := beego.AppConfig.String("upload_path") + filePath

	ok, err := utils.PathExists(path)
	if err != nil {
		logs.Error(c.WorkerId(), " | ", err)
		c.Data["json"] = &resultImage{
			Code:    "fail",
			Message: err.Error(),
		}
		c.ServeJSON()
		return
	}
	if !ok {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			logs.Error(c.WorkerId(), " | ", err)
			c.Data["json"] = &resultImage{
				Code:    "fail",
				Message: err.Error(),
			}
			c.ServeJSON()
			return
		}
	}

	fileName := time.Now().Format("20060102") + "-" + md5FileName + ".jpeg"

	err = c.SaveToFile("MarkdownImage", path+fileName)
	if err != nil {
		logs.Error(c.WorkerId(), " | ", err)
		c.Data["json"] = &resultImage{
			Code:    "fail",
			Message: err.Error(),
		}
		c.ServeJSON()
		return
	}

	image := new(models.Image)
	image.Md5 = md5FileName
	image.RealPath = path + fileName
	image.ImageName = fileName
	image.ClientOriginalName = h.Filename
	image.Extension = ".jpeg"
	image.ImageSize = strconv.Itoa(int(h.Size))
	image.ImagePath = filePath + fileName

	resImg := &resultImage{
		Code: "success",
		Data: struct {
			Width     int    `json:"width"`
			Height    int    `json:"height"`
			Filename  string `json:"filename"`
			Storename string `json:"storename"`
			Size      int64  `json:"size"`
			Path      string `json:"path"`
			Hash      string `json:"hash"`
			Timestamp int64  `json:"timestamp"`
			Url       string `json:"url"`
		}{Width: 0, Height: 0, Filename: h.Filename, Storename: h.Filename, Size: h.Size, Path: filePath, Hash: md5FileName, Timestamp: time.Now().Unix(), Url: beego.AppConfig.String("upload_image_url") + "/images/" + filePath + fileName},
	}

	if _, err := models.AddImage(image); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = resImg
	} else {
		logs.Error(c.WorkerId(), " | ", err)
		c.Data["json"] = &resultImage{
			Code:    "fail",
			Message: err.Error(),
		}
		c.ServeJSON()
		return
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Image
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete/image/:id [delete]
func (c *ImageController) Delete() {
	//idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.ParseInt(idStr, 0, 64)
	//if err := models.DeleteImage(id); err == nil {
	//	c.Data["json"] = "OK"
	//} else {
	//	c.Data["json"] = err.Error()
	//}
	c.ServeJSON()
}
