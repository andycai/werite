package page

import (
	"fmt"
	"html/template"

	"github.com/andycai/werite/administrator/components/user"
	"github.com/andycai/werite/components/page"
	"github.com/andycai/werite/conf"
	"github.com/andycai/werite/constant"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/model"
	"github.com/andycai/werite/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

func handleManagerPage(c *fiber.Ctx) error {
	var (
		total             int64
		totalAll          int64
		totalTrash        int64
		voList            []model.Page
		q                 string
		status            string
		queryParam        string
		CurrentPagination int
	)
	q = c.Query("q")
	status = c.Query("status")

	if c.QueryInt("page") > 1 {
		CurrentPagination = c.QueryInt("page") - 1
	}

	totalAll = page.Dao.Count()
	totalTrash = page.Dao.CountTrash()

	switch status {
	case "trash":
		voList = page.Dao.GetTrashListByPage(CurrentPagination, constant.NUM_PER_PAGE, q)
		total = totalTrash
	default:
		voList = page.Dao.GetListByPage(CurrentPagination, constant.NUM_PER_PAGE, q)
		total = totalAll
	}

	if status != "" {
		queryParam = fmt.Sprintf("&status=%s", status)
	}

	totalPagination, hasPagination := utils.CalcPagination(total)

	bind := fiber.Map{
		"PageTitle":         "All Pages",
		"NavBarActive":      "pages",
		"Path":              "/admin/pages/manager",
		"Pages":             voList,
		"Q":                 q,
		"Status":            status,
		"Total":             total,
		"TotalAll":          totalAll,
		"TotalTrash":        totalTrash,
		"TotalPagination":   totalPagination,
		"HasPagination":     hasPagination,
		"CurrentPagination": CurrentPagination + 1,
		"QueryParam":        template.URL(queryParam),
	}
	user.DecorateUserBar(c, bind)

	return core.Render(c, "admin/pages/pages", bind, "admin/layouts/app")
}

func handleEditorPage(c *fiber.Ctx) error {
	var pageVo model.Page
	hasPage := false

	if c.Params("id") != "" {
		id := cast.ToUint(c.Params("id"))
		hasPage = true
		vo, _ := page.Dao.GetByID(id)
		pageVo = *vo
	}

	bind := fiber.Map{
		"PageTitle":    "Page Editor",
		"NavBarActive": "pages",
		"Path":         "/admin/pages/editor",
		"Domain":       "127.0.0.1",
		"HasPage":      hasPage,
		"Page":         pageVo,
		"MediaPrefix":  conf.GetValue("app.mediaPrefix", "/media/"),
		"MediaHost":    conf.GetValue("app.mediaHost", "http://127.0.0.1:8000"),
	}
	user.DecorateUserBar(c, bind)

	return core.Render(c, "admin/pages/page", bind, "admin/layouts/app")
}

func handleCreate(c *fiber.Ctx) error {
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

func handleUpdate(c *fiber.Ctx) error {
	var pageVo model.Page

	err := page.Bind(c, &pageVo)
	if err != nil {
		return err
	}

	db.Omit("created_at", "user_id").Save(&pageVo)

	core.PushMessages(fmt.Sprintf("Updated page id:%d, title:%s", pageVo.ID, pageVo.Title))

	return c.Redirect("/admin/pages/manager")
}

func handleMoveToTrashByID(c *fiber.Ctx) error {
	id := cast.ToUint(c.Params("id"))
	if id > 0 {
		page.Dao.DeleteByIds([]uint{id})
		core.PushMessages(fmt.Sprintf("Move to trash: %v", id))
	}

	return c.Redirect("/admin/pages/manager")
}

func handleMoveToTrash(c *fiber.Ctx) error {
	form := &utils.FormIDArray{}
	if err := c.BodyParser(form); err != nil {
		return err
	}
	if len(form.ID) > 0 {
		page.Dao.DeleteByIds(form.ID)
		core.PushMessages(fmt.Sprintf("Move to trash: %v", form.ID))
	}

	return c.Redirect("/admin/pages/manager")
}

func DeletePermanetly(c *fiber.Ctx) error {
	form := &utils.FormIDArray{}
	if err := c.BodyParser(form); err != nil {
		return err
	}
	if len(form.ID) > 0 {
		page.Dao.DeletePermanetlyByIds(form.ID)
		core.PushMessages(fmt.Sprintf("Delete permanently: %v", form.ID))
	}

	return c.Redirect("/admin/pages/manager")
}

func handleRestoreByID(c *fiber.Ctx) error {
	id := cast.ToUint(c.Params("id"))
	if id > 0 {
		page.Dao.RestoreByIds([]uint{id})
		core.PushMessages(fmt.Sprintf("Restore pages: %v", id))
	}

	return c.Redirect("/admin/pages/manager")
}

func handleRestore(c *fiber.Ctx) error {
	form := &utils.FormIDArray{}
	if err := c.BodyParser(form); err != nil {
		return err
	}
	if len(form.ID) > 0 {
		page.Dao.RestoreByIds(form.ID)
		core.PushMessages(fmt.Sprintf("Restore pages: %v", form.ID))
	}

	return c.Redirect("/admin/pages/manager")
}
