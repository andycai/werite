package post

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"

	"github.com/andycai/werite/administrator/components/user"
	"github.com/andycai/werite/components/post"
	"github.com/andycai/werite/conf"
	"github.com/andycai/werite/constant"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/model"
	"github.com/andycai/werite/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func handleManagerPage(c *fiber.Ctx) error {
	var (
		total              int64
		totalAll           int64
		totalPublished     int64
		totalDraft         int64
		totalTrash         int64
		voList             []model.Post
		q                  string
		categoryID         int
		status             string
		queryParam         string
		categories         []model.Category
		currrentPagination int
	)
	q = c.Query("q")
	categoryID = c.QueryInt("filter_category") // category
	status = c.Query("status")                 // status: publish, draft or trash

	if c.QueryInt("page") > 1 {
		currrentPagination = c.QueryInt("page") - 1
	}

	totalAll = post.Dao.Count()
	totalPublished = post.Dao.CountByPublished()
	totalDraft = post.Dao.CountByDraft()
	totalTrash = post.Dao.CountByTrash()

	switch status {
	case "publish": // publish
		voList = post.Dao.GetPublishedListByPage(currrentPagination, constant.NUM_PER_PAGE, categoryID, q)
		total = totalPublished
	case "draft": // draft
		voList = post.Dao.GetDraftListByPage(currrentPagination, constant.NUM_PER_PAGE, categoryID, q)
		total = totalDraft
	case "trash": // trash
		voList = post.Dao.GetTrashListByPage(currrentPagination, constant.NUM_PER_PAGE, categoryID, q)
		total = totalTrash
	default: // all
		voList = post.Dao.GetListByPage(currrentPagination, constant.NUM_PER_PAGE, categoryID, q)
		total = totalAll
	}

	if status != "" {
		queryParam = fmt.Sprintf("&status=%s", status)
	}
	if categoryID > 0 {
		queryParam += fmt.Sprintf("&filter_category=%d", categoryID)
	}

	totalPagination, hasPagination := utils.CalcPagination(total)

	categories = post.Dao.GetCategories()
	bind := fiber.Map{
		"PageTitle":         "All Posts",
		"NavBarActive":      "posts",
		"Path":              "/admin/posts/manager",
		"Posts":             voList,
		"Categories":        categories,
		"Q":                 q,
		"Status":            status,
		"Total":             total,
		"TotalAll":          totalAll,
		"TotalPublished":    totalPublished,
		"TotalDraft":        totalDraft,
		"TotalTrash":        totalTrash,
		"FilterCategory":    categoryID,
		"TotalPagination":   totalPagination,
		"HasPagination":     hasPagination,
		"CurrentPagination": currrentPagination + 1,
		"QueryParam":        template.URL(queryParam),
	}
	user.DecorateUserBar(c, bind)

	return core.Render(c, "admin/posts/posts", bind, "admin/layouts/app")
}

func handleEditorPage(c *fiber.Ctx) error {
	var postVo model.Post
	hasPost := false

	if c.Params("id") != "" {
		id := cast.ToUint(c.Params("id"))
		hasPost = true
		vo, _ := post.Dao.GetByID(id)
		postVo = *vo
	}

	categories := post.Dao.GetCategories()
	bind := fiber.Map{
		"PageTitle":    "Post Editor",
		"NavBarActive": "posts",
		"Path":         "/admin/posts/editor",
		"Domain":       "127.0.0.1",
		"HasTags":      true,
		"HasPost":      hasPost,
		"Post":         postVo,
		"Categories":   categories,
		"MediaPrefix":  conf.GetValue("app.mediaPrefix", "/media/"),
		"MediaHost":    conf.GetValue("app.mediaHost", "http://127.0.0.1:8000"),
	}
	user.DecorateUserBar(c, bind)

	return core.Render(c, "admin/posts/post", bind, "admin/layouts/app")
}

func handleCreate(c *fiber.Ctx) error {
	type TagItem struct {
		Value string
	}

	var (
		postVo   model.Post
		tagItems []TagItem
	)

	err := post.Bind(c, &postVo)
	if err != nil {
		return err
	}

	_, userID := authentication.AuthGet(c)
	postVo.UserID = userID

	db.Create(&postVo)

	if c.FormValue("tags") != "" {
		json.Unmarshal([]byte(c.FormValue("tags")), &tagItems)

		for i := 0; i < len(tagItems); i++ {
			tagItem := tagItems[i]
			tag := model.Tag{Name: tagItem.Value}
			tag.Slug = slug.Make(tagItem.Value)

			err := db.Model(&tag).Where("name = ?", tagItem.Value).First(&tag).Error
			if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
				db.Create(&tag)
			}

			if err := db.Model(&postVo).Association("Tags").Append(&tag); err != nil {
				return err
			}
		}
	}

	core.PushMessages(fmt.Sprintf("Created post id:%d, title:%s", postVo.ID, postVo.Title))

	return c.Redirect("/admin/posts/manager")
}

func handleUpdate(c *fiber.Ctx) error {
	type TagItem struct {
		Value string
	}

	var (
		postVo   model.Post
		tagItems []TagItem
		tags     []model.Tag
	)

	err := db.Model(&postVo).
		Where("id = ?", cast.ToUint(c.Params("id"))).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tags.name asc")
		}).
		Find(&postVo).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	err = post.Bind(c, &postVo)
	if err != nil {
		return err
	}

	db.Omit("created_at", "user_id").Save(&postVo)

	if c.FormValue("tags") != "" {
		json.Unmarshal([]byte(c.FormValue("tags")), &tagItems)

		for i := 0; i < len(tagItems); i++ {
			tagItem := tagItems[i]
			tag := model.Tag{Name: tagItem.Value}
			tag.Slug = slug.Make(tagItem.Value)

			err := db.Model(&tag).Where("name = ?", tagItem.Value).First(&tag).Error
			if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
				db.Create(&tag)
			}

			tags = append(tags, tag)
		}

		if err := db.Model(&postVo).Association("Tags").Replace(&tags); err != nil {
			return err
		}
	}

	core.PushMessages(fmt.Sprintf("Updated post id:%d, title:%s", postVo.ID, postVo.Title))

	return c.Redirect("/admin/posts/manager")
}

func handleMoveToTrashByID(c *fiber.Ctx) error {
	id := cast.ToUint(c.Params("id"))
	if id > 0 {
		post.Dao.DeleteByIds([]uint{id})
		core.PushMessages(fmt.Sprintf("Move to trash: %v", id))
	}

	return c.Redirect("/admin/posts/manager")
}

func handleMoveToTrash(c *fiber.Ctx) error {
	form := &utils.FormIDArray{}
	if err := c.BodyParser(form); err != nil {
		return err
	}
	if len(form.ID) > 0 {
		post.Dao.DeleteByIds(form.ID)
		core.PushMessages(fmt.Sprintf("Move to trash: %v", form.ID))
	}

	return c.Redirect("/admin/posts/manager")
}

func DeletePermanetlyByIds(c *fiber.Ctx) error {
	form := &utils.FormIDArray{}
	if err := c.BodyParser(form); err != nil {
		return err
	}
	if len(form.ID) > 0 {
		post.Dao.DeletePermanetlyByIds(form.ID)
		core.PushMessages(fmt.Sprintf("Delete permanetly: %v", form.ID))
	}

	return c.Redirect("/admin/posts/manager")
}

func handleRestoreByID(c *fiber.Ctx) error {
	id := cast.ToUint(c.Params("id"))
	if id > 0 {
		post.Dao.RestoreByIds([]uint{id})
		core.PushMessages(fmt.Sprintf("Restore posts: %v", id))
	}

	return c.Redirect("/admin/posts/manager")
}

func handleRestore(c *fiber.Ctx) error {
	form := &utils.FormIDArray{}
	if err := c.BodyParser(form); err != nil {
		return err
	}
	if len(form.ID) > 0 {
		post.Dao.RestoreByIds(form.ID)
		core.PushMessages(fmt.Sprintf("Restore posts: %v", form.ID))
	}

	return c.Redirect("/admin/posts/manager")
}

func handleManagerCategoryPage(c *fiber.Ctx) error {
	var (
		total             int64
		currentPagination int
		categories        []model.Category
	)
	if c.QueryInt("page") > 1 {
		currentPagination = c.QueryInt("page") - 1
	}

	categories = post.Dao.GetCategoriesByPage(currentPagination, constant.NUM_PER_PAGE)

	total = post.Dao.CountCatgegory()
	totalPagination, hasPagination := utils.CalcPagination(total)

	bind := fiber.Map{
		"PageTitle":         "All Categories",
		"NavBarActive":      "categories",
		"Path":              "/admin/categories/manager",
		"Categories":        categories,
		"TotalPagination":   totalPagination,
		"HasPagination":     hasPagination,
		"CurrentPagination": currentPagination + 1,
	}
	user.DecorateUserBar(c, bind)

	return core.Render(c, "admin/posts/categories", bind, "admin/layouts/app")
}

func handleEditorCategoryPage(c *fiber.Ctx) error {
	var categoryVo model.Category
	hasCategory := false

	if c.Params("id") != "" {
		id := cast.ToUint(c.Params("id"))
		hasCategory = true
		vo, _ := post.Dao.GetCategoryByID(id)
		categoryVo = *vo
	}

	bind := fiber.Map{
		"PageTitle":    "Category Editor",
		"NavBarActive": "categories",
		"Path":         "/admin/categories/editor",
		"Domain":       "127.0.0.1",
		"HasCategory":  hasCategory,
		"Category":     categoryVo,
	}
	user.DecorateUserBar(c, bind)

	return core.Render(c, "admin/posts/category", bind, "admin/layouts/app")
}

func handleCreateCategory(c *fiber.Ctx) error {
	var categoryVo model.Category

	err := post.BindCategory(c, &categoryVo)
	if err != nil {
		return err
	}

	db.Create(&categoryVo)

	core.PushMessages(fmt.Sprintf("Created category id:%d, name:%s", categoryVo.ID, categoryVo.Name))

	return c.Redirect("/admin/categories/manager")
}

func handleUpdateCategory(c *fiber.Ctx) error {
	var categoryVo model.Category

	err := db.Model(&categoryVo).
		Where("id = ?", cast.ToUint(c.Params("id"))).
		Find(&categoryVo).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	err = post.BindCategory(c, &categoryVo)
	if err != nil {
		return err
	}

	db.Save(&categoryVo)

	core.PushMessages(fmt.Sprintf("Updated category id:%d, name:%s", categoryVo.ID, categoryVo.Name))

	return c.Redirect("/admin/categories/manager")
}

func handleDeleteCategories(c *fiber.Ctx) error {
	form := &utils.FormIDArray{}
	if err := c.BodyParser(form); err != nil {
		return err
	}
	if len(form.ID) > 0 {
		post.Dao.DeleteCategoriesByIds(form.ID)
		core.PushMessages(fmt.Sprintf("Delete categories: %v", form.ID))
	}

	return c.Redirect("/admin/categories/manager")
}

func handleManagerTagsPage(c *fiber.Ctx) error {
	var (
		total             int64
		currentPagination int
		tags              []model.Tag
	)
	if c.QueryInt("page") > 1 {
		currentPagination = c.QueryInt("page") - 1
	}

	tags = post.Dao.GetTagsByPage(currentPagination, constant.NUM_PER_PAGE)

	total = post.Dao.CountTag()
	totalPagination, hasPagination := utils.CalcPagination(total)

	bind := fiber.Map{
		"PageTitle":         "All Tags",
		"NavBarActive":      "tags",
		"Path":              "/admin/tags/manager",
		"Tags":              tags,
		"TotalPagination":   totalPagination,
		"HasPagination":     hasPagination,
		"CurrentPagination": currentPagination + 1,
	}
	user.DecorateUserBar(c, bind)

	return core.Render(c, "admin/posts/tags", bind, "admin/layouts/app")
}

func handleEditorTagPage(c *fiber.Ctx) error {
	var tagVo model.Tag
	hasTag := false

	if c.Params("id") != "" {
		id := cast.ToUint(c.Params("id"))
		hasTag = true
		vo, _ := post.Dao.GetTagByID(id)
		tagVo = *vo
	}

	return core.Render(c, "admin/posts/tag", fiber.Map{
		"PageTitle":    "Tag Editor",
		"NavBarActive": "tags",
		"Path":         "/admin/tags/editor",
		"Domain":       "127.0.0.1",
		"HasTag":       hasTag,
		"Tag":          tagVo,
	}, "admin/layouts/app")
}

func handleCreateTag(c *fiber.Ctx) error {
	var tagVo model.Tag

	err := post.BindTag(c, &tagVo)
	if err != nil {
		return err
	}

	db.Create(&tagVo)

	core.PushMessages(fmt.Sprintf("Created Tag id:%d, name:%s", tagVo.ID, tagVo.Name))

	return c.Redirect("/admin/tags/manager")
}

func handleDeleteTags(c *fiber.Ctx) error {
	form := &utils.FormIDArray{}
	if err := c.BodyParser(form); err != nil {
		return err
	}
	if len(form.ID) > 0 {
		post.Dao.DeleteTagsByIds(form.ID)
		core.PushMessages(fmt.Sprintf("Delete tags: %v", form.ID))
	}

	return c.Redirect("/admin/tags/manager")
}
