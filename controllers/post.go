package controllers

import (
	"github.com/astaxie/beego/logs"
	"strconv"
	"github.com/lattecake/blog/models"
	"github.com/shurcooL/github_flavored_markdown"
)

// PostController operations for Post
type PostController struct {
	BaseController
}

// URLMapping ...
func (c *PostController) URLMapping() {
	c.Mapping("Get", c.GetOne)
	c.Mapping("Get", c.GetAll)
}

// GetOne ...
// @Title GetOne
// @Description get Post by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Post
// @Failure 403 :id is empty
// @router /post/:id [get]
func (c *PostController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		logs.Notice(c.WorkerId(), " | ", err)
		c.Abort("404")
		return
	}

	post, err := models.GetPostById(int64(postId))
	if err != nil {
		logs.Warn(c.WorkerId(), " | ", err)
		c.Abort("404")
		return
	}

	images, err := models.GetImageByPostId(post.Id)
	if err != nil {
		logs.Error(c.WorkerId(), " | ", err)
	}

	//unsafe := blackfriday.Run([]byte(post.Content))
	//html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	html := github_flavored_markdown.Markdown([]byte(post.Content))

	c.Data["image"] = images[0]
	c.Data["post"] = post
	c.Data["content"] = string(html)
	c.TplName = "post/detail.tpl"

	post.ReadNum++
	go models.UpdatePostById(post)
}

// GetAll ...
// @Title GetAll
// @Description get Post
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Post
// @Failure 403
// @router /learn [get]
func (c *PostController) GetAll() {

}
