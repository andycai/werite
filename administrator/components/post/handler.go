package post

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/andycai/werite/administrator/enum"
	"github.com/andycai/werite/administrator/utils"
	"github.com/andycai/werite/components/post"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func ManagerPage(c *fiber.Ctx) error {
	var (
		queryParam string
	)
	q := c.Query("q")
	qc := c.QueryInt("qc") // category
	qf := c.QueryInt("qf") // filter: published:1, draft:2 or trash:3

	currrentPagination := 0
	if c.QueryInt("page") > 1 {
		currrentPagination = c.QueryInt("page") - 1
	}

	voList := post.Dao.GetListByPage(currrentPagination, enum.NUM_PER_PAGE)

	total := post.Dao.Count()
	switch qf {
	case 1:
		total = post.Dao.CountByPublished()
	case 2:
		total = post.Dao.CountByDraft()
	case 3:
		total = post.Dao.CountByTrash()
	}

	if qf > 0 {
		queryParam = fmt.Sprintf("%s&qf=%d", queryParam, qf)
	}
	if qc > 0 {
		queryParam = fmt.Sprintf("%s&qc=%d", queryParam, qc)
	}

	totalPagination, hasPagination := utils.CalcPagination(total)

	categories := post.Dao.GetCategories()
	return core.Render(c, "admin/posts/posts", fiber.Map{
		"PageTitle":         "All Posts",
		"NavBarActive":      "posts",
		"Path":              "/admin/posts/manager",
		"Posts":             voList,
		"Categories":        categories,
		"Total":             total,
		"Q":                 q,
		"QC":                qc,
		"QF":                qf,
		"QueryParam":        queryParam,
		"TotalPagination":   totalPagination,
		"HasPagination":     hasPagination,
		"CurrentPagination": currrentPagination + 1,
	}, "admin/layouts/app")
}

func EditorPage(c *fiber.Ctx) error {
	var postVo model.Post
	hasPost := false

	if c.Params("id") != "" {
		id := cast.ToUint(c.Params("id"))
		hasPost = true
		vo, _ := post.Dao.GetByID(id)
		postVo = *vo
	}

	categories := post.Dao.GetCategories()

	return core.Render(c, "admin/posts/post", fiber.Map{
		"PageTitle":    "Post Editor",
		"NavBarActive": "posts",
		"Path":         "/admin/posts/editor",
		"Domain":       "127.0.0.1",
		"HasPost":      hasPost,
		"Post":         postVo,
		"Categories":   categories,
	}, "admin/layouts/app")
}

func Create(c *fiber.Ctx) error {
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

func Update(c *fiber.Ctx) error {
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

func ManagerCategoryPage(c *fiber.Ctx) error {
	currentPagination := 0
	if c.QueryInt("page") > 1 {
		currentPagination = c.QueryInt("page") - 1
	}

	categories := post.Dao.GetCategoriesByPage(currentPagination, enum.NUM_PER_PAGE)

	totalPagination, hasPagination := utils.CalcPagination(post.Dao.CatgegoryCount())

	return core.Render(c, "admin/posts/categories", fiber.Map{
		"PageTitle":         "All Categories",
		"NavBarActive":      "categories",
		"Path":              "/admin/categories/manager",
		"Categories":        categories,
		"TotalPagination":   totalPagination,
		"HasPagination":     hasPagination,
		"CurrentPagination": currentPagination + 1,
	}, "admin/layouts/app")
}

func EditorCategoryPage(c *fiber.Ctx) error {
	var categoryVo model.Category
	hasCategory := false

	if c.Params("id") != "" {
		id := cast.ToUint(c.Params("id"))
		hasCategory = true
		vo, _ := post.Dao.GetCategoryByID(id)
		categoryVo = *vo
	}

	return core.Render(c, "admin/posts/category", fiber.Map{
		"PageTitle":    "Category Editor",
		"NavBarActive": "categories",
		"Path":         "/admin/categories/editor",
		"Domain":       "127.0.0.1",
		"HasCategory":  hasCategory,
		"Category":     categoryVo,
	}, "admin/layouts/app")
}

func CreateCategory(c *fiber.Ctx) error {
	var categoryVo model.Category

	err := post.BindCategory(c, &categoryVo)
	if err != nil {
		return err
	}

	db.Create(&categoryVo)

	core.PushMessages(fmt.Sprintf("Created category id:%d, name:%s", categoryVo.ID, categoryVo.Name))

	return c.Redirect("/admin/categories/manager")
}

func UpdateCategory(c *fiber.Ctx) error {
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

func ManagerTagsPage(c *fiber.Ctx) error {
	currentPagination := 0
	if c.QueryInt("page") > 1 {
		currentPagination = c.QueryInt("page") - 1
	}

	tags := post.Dao.GetTagsByPage(currentPagination, enum.NUM_PER_PAGE)

	totalPagination, hasPagination := utils.CalcPagination(post.Dao.TagCount())

	return core.Render(c, "admin/posts/tags", fiber.Map{
		"PageTitle":         "All Tags",
		"NavBarActive":      "tags",
		"Path":              "/admin/tags/manager",
		"Tags":              tags,
		"TotalPagination":   totalPagination,
		"HasPagination":     hasPagination,
		"CurrentPagination": currentPagination + 1,
	}, "admin/layouts/app")
}

func EditorTagPage(c *fiber.Ctx) error {
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

func CreateTag(c *fiber.Ctx) error {
	var tagVo model.Tag

	err := post.BindTag(c, &tagVo)
	if err != nil {
		return err
	}

	db.Create(&tagVo)

	core.PushMessages(fmt.Sprintf("Created Tag id:%d, name:%s", tagVo.ID, tagVo.Name))

	return c.Redirect("/admin/tags/manager")
}
