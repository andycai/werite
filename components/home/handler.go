package page

import (
	"math"

	"github.com/andycai/werite/components/post"
	mpo "github.com/andycai/werite/components/post/model"
	"github.com/andycai/werite/components/user"
	mu "github.com/andycai/werite/components/user/model"
	"github.com/andycai/werite/core"
	"github.com/andycai/werite/library/authentication"
	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	var authenticatedUser *mu.User

	isAuthenticated, userID := authentication.AuthGet(c)
	if isAuthenticated {
		authenticatedUser = user.Dao.GetByID(userID)
	}

	return core.Render(c, "home/index", fiber.Map{
		"PageTitle":         "Home Page - Werite",
		"FiberCtx":          c,
		"NavBarActive":      "home",
		"AuthenticatedUser": authenticatedUser,
		"CurrentPage":       c.QueryInt("page"),
	}, "layouts/app")
}

//#region HTMX interface

// HTMXHomePage home page
func HTMXHomePage(c *fiber.Ctx) error {
	var authenticatedUser *mu.User

	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		authenticatedUser = user.Dao.GetByID(userID)
	}

	return core.Render(c, "home/htmx-home-page", fiber.Map{
		"PageTitle":         "Home",
		"NavBarActive":      "home",
		"FiberCtx":          c,
		"AuthenticatedUser": authenticatedUser,
	}, "layouts/app-htmx")
}

// HTMXHomeTagList tag list
func HTMXHomeTagList(c *fiber.Ctx) error {
	var (
		tags    []mpo.Tag
		hasTags bool
	)
	return core.Render(c, "home/partials/tag-item-list", fiber.Map{
		"Tags":    tags,
		"HasTags": hasTags,
	}, "layouts/app-htmx")
}

// HTMXHomeGlobalFeed global feed
func HTMXHomeGlobalFeed(c *fiber.Ctx) error {
	var (
		posts           []mpo.Post
		hasPosts        bool
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

	count = post.Dao.Count()
	posts = post.Dao.GetListByPage(page, numPerPage)

	feedNavbarItems := []fiber.Map{
		{
			"Title":     "Global Feed",
			"IsActive":  true,
			"HXPushURL": "/",
			"HXGetURL":  "/htmx/home/global-feed",
		},
	}

	if count > 0 && (count/int64(numPerPage) > 0) {
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

	if len(posts) > 0 {
		hasPosts = true
	}

	return core.Render(c, "home/htmx-home-feed", fiber.Map{
		"HasPosts":            hasPosts,
		"Posts":               posts,
		"FeedNavbarItems":     feedNavbarItems,
		"AuthenticatedUserID": userID,
		"TotalPagination":     totalPagination,
		"HasPagination":       hasPagination,
		"CurrentPagination":   page + 1,
		"PathPagination":      "global-feed",
	}, "layouts/app-htmx")
}

//#endregion
