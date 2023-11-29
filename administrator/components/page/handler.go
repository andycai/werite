package page

import (
	"fmt"
	"math"

	"github.com/andycai/werite/components/page"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/model"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

func ManagerPage(c *fiber.Ctx) error {
	var (
		hasPagination   bool
		totalPagination int
		count           int64
	)
	q := c.Query("q")
	qc := c.QueryInt("qc")

	numPerPage := 10
	curPage := 0
	if c.QueryInt("page") > 1 {
		curPage = c.QueryInt("page") - 1
	}

	count = page.Dao.Count()
	voList := page.Dao.GetListByPage(curPage, numPerPage)

	if count > 0 && (count/int64(numPerPage) > 0) {
		pageDivision := float64(count) / float64(numPerPage)
		totalPagination = int(math.Ceil(pageDivision))
		hasPagination = true
	}

	return core.Render(c, "admin/pages/pages", fiber.Map{
		"PageTitle":         "All Pages",
		"NavBarActive":      "pages",
		"Path":              "/admin/pages/manager",
		"Pages":             voList,
		"Q":                 q,
		"QC":                qc,
		"TotalPagination":   totalPagination,
		"HasPagination":     hasPagination,
		"CurrentPagination": curPage + 1,
	}, "admin/layouts/app")
}

func EditorPage(c *fiber.Ctx) error {
	var pageVo model.Page
	hasPage := false

	if c.Params("id") != "" {
		id := cast.ToUint(c.Params("id"))
		hasPage = true
		vo, _ := page.Dao.GetByID(id)
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
	pageVo.UserID = userID

	db.Create(&pageVo)

	core.PushMessages(fmt.Sprintf("Created page id:%d, title:%s", pageVo.ID, pageVo.Title))

	return c.Redirect("/admin/pages/manager")
}

func Update(c *fiber.Ctx) error {
	var pageVo model.Page

	err := page.Bind(c, &pageVo)
	if err != nil {
		return err
	}

	db.Omit("created_at", "user_id").Save(&pageVo)

	core.PushMessages(fmt.Sprintf("Updated page id:%d, title:%s", pageVo.ID, pageVo.Title))

	return c.Redirect("/admin/pages/manager")
}
