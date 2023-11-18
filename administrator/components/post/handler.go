package post

import (
	"github.com/andycai/werite/components/post"
	"github.com/andycai/werite/components/post/model"
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
)

func PostsPage(c *fiber.Ctx) error {
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
		"Path":         "/admin/posts",
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

func PostPage(c *fiber.Ctx) error {
	return core.Render(c, "admin/posts/post", fiber.Map{
		"PageTitle":    "Post",
		"NavBarActive": "posts",
		"Path":         "/admin/post",
		"Domain":       "127.0.0.1",
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
		"Path":         "/admin/post",
		"Domain":       "127.0.0.1",
		"HasPost":      hasPost,
		"Post":         postVo,
	}, "admin/layouts/app")
}

func EditorAction(c *fiber.Ctx) error {
	return c.Redirect("/admin/posts")
}
