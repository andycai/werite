package handler

import (
	"math"

	"github.com/andycai/werite/library/authentication"
	"github.com/andycai/werite/v2/model"
	"github.com/andycai/werite/v2/system"
	"github.com/gofiber/fiber/v2"
)

func HomePage(c *Ctx) error {
	var authenticatedUser *model.User

	isAuthenticated, userID := authentication.AuthGet(c)
	if isAuthenticated {
		authenticatedUser = system.User.GetByID(userID)
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
	var authenticatedUser *model.User

	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		authenticatedUser = system.User.GetByID(userID)
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

	count = system.Article.Count()
	articles = system.Article.GetListByPage(page, numPerPage)

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
