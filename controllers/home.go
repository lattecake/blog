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

// home ...
// @Title Create
// @Description create Home
// @Param	body		body 	models.Home	true		"body for Home content"
// @Success 201 {object} models.Home
// @Failure 403 body is empty
// @router / [get]
func (c *HomeController) Index() {

	// TODO 由于服务器本身就是单核，所以就不用多线程处理了，没啥效果

	posts, err := models.GetAllPost(map[string]string{
		"status": "1",
	}, []string{
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

	query := make(map[string]string)
	query["Star"] = "1"

	topPosts, err := models.GetAllPost(query, []string{
		"Title", "Description", "Id", "PushTime", "Star",
	}, []string{
		"PushTime",
	}, []string{
		"desc",
	}, 0, 3)

	var postIds []int64

	if posts != nil {
		for _, val := range posts {
			postIds = append(postIds, val.(map[string]interface{})["Id"].(int64))
		}
		for _, val := range topPosts {
			postIds = append(postIds, val.(map[string]interface{})["Id"].(int64))
		}
	}

	postImage := map[int64]string{}

	if postIds != nil {
		images, err := models.GetImageByPostIds(postIds)
		if err != nil {
			logs.Error(c.WorkerId(), " | ", err)
			c.Abort("500")
			return
		}
		for _, image := range images {
			postImage[image.PostId] = image.RealPath
		}

		for k, v := range posts {
			posts[k].(map[string]interface{})["Image"] = postImage[v.(map[string]interface{})["Id"].(int64)]
		}
	}

	if err != nil {
		logs.Error(" | ", c.WorkerId(), " | ", err)
	}

	for k, v := range topPosts {
		topPosts[k].(map[string]interface{})["Image"] = postImage[v.(map[string]interface{})["Id"].(int64)]
	}

	c.Data["postImages"] = postImage
	c.Data["starPosts"] = topPosts
	c.Data["posts"] = posts
	c.TplName = "home/index.tpl"
}

// reward ...
// @Title Create
// @Description create Home
// @Param	body		body 	models.Home	true		"body for Home content"
// @Success 201 {object} models.Home
// @Failure 403 body is empty
// @router /reward [get]
func (c *HomeController) Reward() {

	c.Data["title"] = "谢谢老板们打赏"
	c.TplName = "home/reward.tpl"
}

// About ...
// @Title Create
// @Description create Home
// @Param	body		body 	models.Home	true		"body for Home content"
// @Success 201 {object} models.Home
// @Failure 403 body is empty
// @router /about [get]
func (c *HomeController) About() {

	c.Data["title"] = "关于我的都在这里"
	c.TplName = "home/about.tpl"
}
