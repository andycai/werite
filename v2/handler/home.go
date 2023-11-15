package handler

import (
	"math"

	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/library/database"
	"github.com/andycai/werite/v2/dao"
	"github.com/andycai/werite/v2/model"
	"github.com/gofiber/fiber/v2"
)

func HomePage(c *Ctx) error {
	var authenticatedUser model.User

	isAuthenticated, userID := authentication.AuthGet(c)
	if isAuthenticated {
		db := database.Get()
		db.Model(&authenticatedUser).
			Where("id = ?", userID).
			First(&authenticatedUser)
	}

	return Render(c, "home/index", fiber.Map{
		"PageTitle":         "Home Page - Werite",
		"FiberCtx":          c,
		"NavBarActive":      "home",
		"AuthenticatedUser": authenticatedUser,
		"CurrentPage":       c.QueryInt("page"),
	}, "layouts/app")
}

//#region HTMX interface

// HTMXHomePage home page
func HTMXHomePage(c *Ctx) error {
	var authenticatedUser model.User

	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		db := database.Get()
		db.Model(&authenticatedUser).
			Where("id = ?", userID).
			First(&authenticatedUser)
	}

	return Render(c, "home/htmx-home-page", fiber.Map{
		"PageTitle":         "Home",
		"NavBarActive":      "home",
		"FiberCtx":          c,
		"AuthenticatedUser": authenticatedUser,
	}, "layouts/app-htmx")
}

// HTMXHomeTagList tag list
func HTMXHomeTagList(c *Ctx) error {
	var (
		tags    []model.Tag
		hasTags bool
	)
	return Render(c, "home/partials/tag-item-list", fiber.Map{
		"Tags":    tags,
		"HasTags": hasTags,
	}, "layouts/app-htmx")
}

// HTMXHomeGlobalFeed global feed
func HTMXHomeGlobalFeed(c *Ctx) error {
	var (
		articles        []model.Article
		hasArticles     bool
		hasPagination   bool
		totalPagination int
		count           int64
	)

	numPerPage := 5
	page := 0
	if c.QueryInt("page") > 1 {
		page = c.QueryInt("page") - 1
	}

	isAuthenticated, userID := authentication.AuthGet(c)

	a := dao.Article
	count, _ = a.Count()
	a.Offset(page * numPerPage).Limit(numPerPage).Order(a.CreatedAt.Desc()).Scan(&articles)

	// db := database.Get()
	// db.Debug().Model(&articles).
	// 	// Preload("Tags", func(db *gorm.DB) *gorm.DB {
	// 	// return db.Order("tags.name asc")
	// 	// }).
	// 	Limit(numPerPage).
	// 	Offset(page * numPerPage).
	// 	Order("created_at desc").
	// 	Find(&articles)

	// db.Model(&articles).Count(&count)

	feedNavbarItems := []fiber.Map{
		{
			"Title":     "Global Feed",
			"IsActive":  true,
			"HXPushURL": "/",
			"HXGetURL":  "/htmx/home/global-feed",
		},
	}

	if count > 0 && (count/5 > 0) {
		pageDivision := float64(count) / float64(5)
		totalPagination = int(math.Ceil(pageDivision))
		hasPagination = true
	}

	if isAuthenticated {

		feedNavbarItems = append([]fiber.Map{
			{
				"Title":     "Your Feed",
				"IsActive":  false,
				"HXPushURL": "/your-feed",
				"HXGetURL":  "/htmx/home/your-feed",
			},
		}, feedNavbarItems...)
	}

	if len(articles) > 0 {
		hasArticles = true
	}

	return Render(c, "home/htmx-home-feed", fiber.Map{
		"HasArticles":         hasArticles,
		"Articles":            articles,
		"FeedNavbarItems":     feedNavbarItems,
		"AuthenticatedUserID": userID,
		"TotalPagination":     totalPagination,
		"HasPagination":       hasPagination,
		"CurrentPagination":   page + 1,
		"PathPagination":      "global-feed",
	}, "layouts/app-htmx")
}

//#endregion
