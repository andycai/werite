package page

import (
	"github.com/andycai/werite/components/page"
	"github.com/andycai/werite/components/page/model"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/gofiber/fiber/v2"
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

	// count = page.Dao.Count()
	voList := page.Dao.GetListByPage(curPage, numPerPage)

	return core.Render(c, "admin/pages/pages", fiber.Map{
		"PageTitle":    "All Pages",
		"NavBarActive": "pages",
		"Path":         "/admin/pages/manager",
		"Pages":        voList,
		"Q":            q,
		"QC":           qc,
		"PG":           curPage,
		"Prev":         0,
		"Next":         0,
		"PP":           map[int]string{},
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
		"Path":         "/admin/pages/editor",
		"Domain":       "127.0.0.1",
		"HasPage":      hasPage,
		"Page":         pageVo,
	}, "admin/layouts/app")
}

func Create(c *fiber.Ctx) error {
	var pageVo model.Page

	err := page.Bind(c, &pageVo)
	if err != nil {
		return err
	}

	_, userID := authentication.AuthGet(c)
	pageVo.UserID = cast.ToInt32(userID)

	db.Create(&pageVo)

	return c.Redirect("/admin/pages/manager")
}

func Update(c *fiber.Ctx) error {
	var pageVo model.Page

	err := page.Bind(c, &pageVo)
	if err != nil {
		return err
	}

	db.Omit("created_at", "user_id").Save(pageVo)

	return c.Redirect("/admin/pages/manager")
}
