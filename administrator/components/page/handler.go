package page

import (
	"github.com/andycai/werite/components/page"
	"github.com/andycai/werite/components/page/model"
	"github.com/andycai/werite/core"
	"github.com/gofiber/fiber/v2"
)

func PagesPage(c *fiber.Ctx) error {
	q := c.Query("q")
	qc := c.QueryInt("qc")

	numPerPage := 10
	curPage := 0
	if c.QueryInt("page") > 1 {
		curPage = c.QueryInt("page") - 1
	}

	// count = page.Dao.Count()
	voList := page.Dao.GetListByPage(curPage, numPerPage)

	return core.Render(c, "admin/pages/pages", fiber.Map{
		"PageTitle":    "All Pages",
		"NavBarActive": "pages",
		"Path":         "/admin/pages",
		"Pages":        voList,
		"Q":            q,
		"QC":           qc,
		"PG":           curPage,
		"Prev":         0,
		"Next":         0,
		"PP":           map[int]string{},
	}, "admin/layouts/app")
}

func PagePage(c *fiber.Ctx) error {
	return core.Render(c, "admin/pages/page", fiber.Map{
		"PageTitle":    "Page",
		"NavBarActive": "pages",
		"Path":         "/admin/page",
		"Domain":       "127.0.0.1",
	}, "admin/layouts/app")
}

func EditorPage(c *fiber.Ctx) error {
	var pageVo model.Page
	hasPage := false

	if c.Params("slug") != "" {
		slug := c.Params("slug")
		hasPage = true
		vo, _ := page.Dao.GetBySlug(slug)
		pageVo = *vo
	}

	return core.Render(c, "admin/pages/page", fiber.Map{
		"PageTitle":    "Page Editor",
		"NavBarActive": "pages",
		"Path":         "/admin/page",
		"Domain":       "127.0.0.1",
		"HasPage":      hasPage,
		"Page":         pageVo,
	}, "admin/layouts/app")
}

func PageAction(c *fiber.Ctx) error {
	return c.Redirect("admin/pages")
}
