package controllers

import (
	"github.com/lattecake/blog/models"
	"github.com/astaxie/beego/logs"
)

// HomeController operations for Home
type HomeController struct {
	BaseController
}

// URLMapping ...
func (c *HomeController) URLMapping() {
	c.Mapping("Get", c.Index)
}

// Post ...
// @Title Create
// @Description create Home
// @Param	body		body 	models.Home	true		"body for Home content"
// @Success 201 {object} models.Home
// @Failure 403 body is empty
// @router /:pageId [get]
func (c *HomeController) Index() {

	posts, err := models.GetAllPost(map[string]string{}, []string{
		"Title", "Description", "Id", "PushTime",
	}, []string{
		"Id",
	}, []string{
		"desc",
	}, 0, 10)
	if err != nil {
		logs.Critical(c.WorkerId(), " | ", err)
		c.Abort("500")
		return
	}

	var postIds []int64

	for _, val := range posts {
		postIds = append(postIds, val.(map[string]interface{})["Id"].(int64))
	}

	images, err := models.GetImageByPostIds(postIds)
	if err != nil {
		logs.Error(c.WorkerId(), " | ", err)
		c.Abort("500")
		return
	}

	postImage := map[int64]string{}

	for _, image := range images {
		postImage[image.PostId] = image.RealPath
	}

	logs.Warn(postImage)

	c.Data["postImages"] = postImage
	c.Data["posts"] = posts
	c.TplName = "home/index.tpl"
}
