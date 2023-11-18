package page

import (
	"github.com/andycai/werite/components/page"
	"github.com/andycai/werite/components/page/model"
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

func StorePage(c *fiber.Ctx) error {
	id := cast.ToInt32(c.FormValue("id"))
	slugVal := c.FormValue("slug")
	title := c.FormValue("title")

	if slugVal == "" {
		slugVal = slug.Make(title)
	}

	content := c.FormValue("content")
	publishAt := core.ParseDate(c.FormValue("publish_at"))

	page := &model.Page{
		ID:        id,
		Slug:      slugVal,
		Title:     title,
		Body:      content,
		PublishAt: publishAt,
	}

	err := core.Validate(page)
	if err != nil {
		return err
	}

	db.Create(page)

	return c.Redirect("admin/pages")
}

func UpdatePage(c *fiber.Ctx) error {
	id := cast.ToInt32(c.FormValue("id"))
	slugVal := c.FormValue("slug")
	title := c.FormValue("title")

	if slugVal == "" {
		slugVal = slug.Make(title)
	}

	content := c.FormValue("content")
	publishAt := core.ParseDate(c.FormValue("publish_at"))

	page := &model.Page{
		ID:        id,
		Slug:      slugVal,
		Title:     title,
		Body:      content,
		PublishAt: publishAt,
	}

	err := core.Validate(page)
	if err != nil {
		return err
	}

	db.Save(page)

	return c.Redirect("admin/pages")
}
