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

	c.Data["title"] = post.Title
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

	pageId, err := strconv.Atoi(c.Ctx.Input.Param(":pageId"))
	action := c.GetString("action", "all")

	if err != nil {
		logs.Warn(" | ", c.WorkerId(), " | ", err)
		pageId = 1
	}

	if pageId < 1 {
		pageId = 1
	}

	var offset int64 = 0
	var limit int64 = 10

	if pageId > 1 {
		offset = int64(pageId) * limit
	}

	query := make(map[string]string)
	query["status"] = "1"

	switch action {
	case "learn":
		query["action__in"] = "1"
	case "life":
		query["action__in"] = "2,3"
	}

	posts, err := models.GetAllPost(query, []string{
		"Title", "Description", "Id", "PushTime",
	}, []string{
		"Id",
	}, []string{
		"desc",
	}, offset, limit)
	if err != nil {
		logs.Critical(c.WorkerId(), " | ", err)
		c.Abort("500")
		return
	}

	var postIds []int64

	if posts != nil {
		for _, val := range posts {
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

	c.Data["title"] = "小小的记录"
	c.Data["postImages"] = postImage
	c.Data["posts"] = posts
	c.Data["prev"] = pageId - 1
	c.Data["next"] = pageId + 1

	c.TplName = "post/list.tpl"
}
