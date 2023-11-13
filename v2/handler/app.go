package handler

import (
	"errors"

	"github.com/andycai/werite/library/authentication"
	database "github.com/andycai/werite/library/database/gorm"
	"github.com/andycai/werite/v2/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppHandler struct{}

var App = new(AppHandler)

//#region private methods

//

//#endregion

//#region handlers

// HomePage render home page
func (ah AppHandler) HomePage(c *Ctx) error {
	var authenticatedUser model.User

	// isAuthenticated, userID := authentication.AuthGet(c)

	return render(c, "home/index", fiber.Map{
		"PageTitle":         "Andy's Blog Home Page",
		"FiberCtx":          c,
		"NavBarActive":      "home",
		"AuthenticatedUser": authenticatedUser,
		"CurrentPage":       c.QueryInt("page"),
	}, "layouts/app")
}

// HomePage render home page
func (ah AppHandler) ArticleDetailPage(c *Ctx) error {
	var article model.Article
	var authenticatedUser model.User
	isSelf := false
	isFollowed := false

	isAuthenticated, userID := authentication.AuthGet(c)

	db := database.Get()

	err := db.Model(&article).
		Where("slug = ?", c.Params("slug")).
		Find(&article).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Redirect("/")
		}
	}

	if isAuthenticated {
		db.Model(&authenticatedUser).
			Where("id = ?", userID).
			First(&authenticatedUser)
	}

	return c.Render("articles/show", fiber.Map{
		"PageTitle":          article.Title + " — Werite",
		"Article":            article,
		"FiberCtx":           c,
		"IsOob":              false,
		"IsSelf":             isSelf,
		"IsFollowed":         isFollowed,
		"IsArticleFavorited": false,
		"AuthenticatedUser":  authenticatedUser,
	}, "layouts/app")
}

// HomePage render home page
func (ah AppHandler) Page(c *Ctx) error {
	return nil
}

// HomePage render home page
func (ah AppHandler) Series(c *Ctx) error {
	return nil
}

// HomePage render home page
func (ah AppHandler) Archive(c *Ctx) error {
	return nil
}

// HomePage render home page
func (ah AppHandler) Search(c *Ctx) error {
	return nil
}

//#endregion
