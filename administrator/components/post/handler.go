package post

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"

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
	q := c.Query("q")
	qc := c.QueryInt("qc")

	numPerPage := 10
	curPage := 0
	if c.QueryInt("page") > 1 {
		curPage = c.QueryInt("page") - 1
	}

	// count = post.Dao.Count()
	voList := post.Dao.GetListByPage(curPage, numPerPage)

	categories := post.Dao.GetCategories()
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
	numPerPage := 10
	curPage := 0
	if c.QueryInt("page") > 1 {
		curPage = c.QueryInt("page") - 1
	}

	categories := post.Dao.GetCategoriesByPage(curPage, numPerPage)

	return core.Render(c, "admin/posts/categories", fiber.Map{
		"PageTitle":    "All Categories",
		"NavBarActive": "categories",
		"Path":         "/admin/categories/manager",
		"Categories":   categories,
		"Page":         curPage,
		"Prev":         0,
		"Next":         0,
		"PP":           map[int]string{},
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
	var (
		hasPagination   bool
		totalPagination int
		count           int64
	)
	numPerPage := 2
	page := 0

	if c.QueryInt("page") > 1 {
		page = c.QueryInt("page") - 1
	}
	count = post.Dao.TagCount()
	tags := post.Dao.GetTagsByPage(page, numPerPage)

	if count > 0 && (count/int64(numPerPage) > 0) {
		pageDivision := float64(count) / float64(numPerPage)
		totalPagination = int(math.Ceil(pageDivision))
		hasPagination = true
	}

	return core.Render(c, "admin/posts/tags", fiber.Map{
		"PageTitle":         "All Tags",
		"NavBarActive":      "tags",
		"Path":              "/admin/tags/manager",
		"Tags":              tags,
		"TotalPagination":   totalPagination,
		"HasPagination":     hasPagination,
		"CurrentPagination": page + 1,
		"PathPagination":    "/admin/tags/manager",
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
