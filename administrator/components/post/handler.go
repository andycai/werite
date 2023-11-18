package post

import (
	"github.com/andycai/werite/components/post"
	"github.com/andycai/werite/components/post/model"
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/spf13/cast"
)

func ManagerPage(c *fiber.Ctx) error {
	q := c.Query("q")
	qc := c.QueryInt("qc")

	numPerPage := 10
	curPage := 0
	if c.QueryInt("page") > 1 {
		curPage = c.QueryInt("page") - 1
	}

	// count = post.Dao.Count()
	voList := post.Dao.GetListByPage(curPage, numPerPage)

	categories := []*model.Category{}
	return core.Render(c, "admin/posts/posts", fiber.Map{
		"PageTitle":    "All Posts",
		"NavBarActive": "posts",
		"Path":         "/admin/posts/manager",
		"Posts":        voList,
		"Categories":   categories,
		"Q":            q,
		"QC":           qc,
		"Page":         curPage,
		"Prev":         0,
		"Next":         0,
		"PP":           map[int]string{},
	}, "admin/layouts/app")
}

func EditorPage(c *fiber.Ctx) error {
	var postVo model.Post
	hasPost := false

	if c.Params("slug") != "" {
		slug := c.Params("slug")
		hasPost = true
		vo, _ := post.Dao.GetBySlug(slug)
		postVo = *vo
	}

	return core.Render(c, "admin/posts/post", fiber.Map{
		"PageTitle":    "Post Editor",
		"NavBarActive": "posts",
		"Path":         "/admin/posts/editor",
		"Domain":       "127.0.0.1",
		"HasPost":      hasPost,
		"Post":         postVo,
	}, "admin/layouts/app")
}

func StorePage(c *fiber.Ctx) error {
	id := cast.ToInt32(c.FormValue("id"))
	slugVal := c.FormValue("slug")
	title := c.FormValue("title")
	description := c.FormValue("description")
	isDraft := cast.ToInt32(c.FormValue("isDraft"))

	if slugVal == "" {
		slugVal = slug.Make(title)
	}

	content := c.FormValue("content")
	publishAt := core.ParseDate(c.FormValue("publish_at"))

	post := &model.Post{
		ID:          id,
		Slug:        slugVal,
		Title:       title,
		Description: description,
		IsDraft:     isDraft,
		Body:        content,
		PublishAt:   publishAt,
	}

	err := core.Validate(post)
	if err != nil {
		return err
	}

	db.Create(post)

	return c.Redirect("admin/posts/manager")
}

func UpdatePage(c *fiber.Ctx) error {
	id := cast.ToInt32(c.FormValue("id"))
	slugVal := c.FormValue("slug")
	title := c.FormValue("title")
	description := c.FormValue("description")

	if slugVal == "" {
		slugVal = slug.Make(title)
	}

	content := c.FormValue("content")
	publishAt := core.ParseDate(c.FormValue("publish_at"))

	post := &model.Post{
		ID:          id,
		Slug:        slugVal,
		Title:       title,
		Description: description,
		Body:        content,
		PublishAt:   publishAt,
	}

	err := core.Validate(post)
	if err != nil {
		return err
	}

	db.Save(post)

	return c.Redirect("admin/posts/manager")
}
