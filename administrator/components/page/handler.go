package page

import (
	"github.com/andycai/werite/components/page"
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
	pages := page.Dao.GetListByPage(curPage, numPerPage)

	return core.Render(c, "admin/pages/pages", fiber.Map{
		"PageTitle":    "All Pages",
		"NavBarActive": "pages",
		"Path":         "/admin/pages",
		"Pages":        pages,
		"Q":            q,
		"QC":           qc,
		"PG":           curPage,
		"Prev":         0,
		"Next":         0,
		"PP":           map[int]string{},
	}, "admin/layouts/app")
}

func PagePage(c *fiber.Ctx) error {
	return nil
}
