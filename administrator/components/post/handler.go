package post

import (
	"encoding/json"
	"errors"

	"github.com/andycai/werite/components/post"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/model"
	"github.com/gofiber/fiber/v2"
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

			err := db.Model(&tag).Where("name = ?", tagItem.Value).First(&tag).Error
			if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
				db.Create(&tag)
			}

			if err := db.Model(&postVo).Association("Tags").Append(&tag); err != nil {
				return err
			}
		}
	}

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
	return nil
}

func CreateCategory(c *fiber.Ctx) error {
	return nil
}

func UpdateCategory(c *fiber.Ctx) error {
	return nil
}
