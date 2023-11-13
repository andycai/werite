package handler

import (
	"errors"
	"math"

	"github.com/andycai/werite/library/authentication"
	database "github.com/andycai/werite/library/database/gorm"
	"github.com/andycai/werite/v2/dao"
	"github.com/andycai/werite/v2/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HTMXHandler struct{}

var HTMX = new(HTMXHandler)

//#region private methods

//

//#endregion

//#region handler

// HomeTagList tag list
func (hh HTMXHandler) HomePage(c *Ctx) error {
	var authenticatedUser model.User

	isAuthenticated, userID := authentication.AuthGet(c)

	if isAuthenticated {
		db := database.Get()
		db.Model(&authenticatedUser).
			Where("id = ?", userID).
			First(&authenticatedUser)
	}

	return render(c, "home/htmx-home-page", fiber.Map{
		"PageTitle":         "Home",
		"NavBarActive":      "home",
		"FiberCtx":          c,
		"AuthenticatedUser": authenticatedUser,
	}, "layouts/app-htmx")
}

// HomeTagList tag list
func (hh HTMXHandler) HomeTagList(c *Ctx) error {
	var (
		tags    []model.Tag
		hasTags bool
	)
	return render(c, "home/partials/tag-item-list", fiber.Map{
		"Tags":    tags,
		"HasTags": hasTags,
	}, "layouts/app-htmx")
}

// HomeGlobalFeed global feed
func (hh HTMXHandler) HomeGlobalFeed(c *Ctx) error {
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

	return render(c, "home/htmx-home-feed", fiber.Map{
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

// HomeArticleDetailPage detail page
func (hh HTMXHandler) HomeArticleDetailPage(c *Ctx) error {
	var article model.Article
	isSelf := false
	isFollowed := false
	var authenticatedUser model.User

	isAuthenticated, userID := authentication.AuthGet(c)

	db := database.Get()

	if isAuthenticated {
		db.Model(&authenticatedUser).
			Where("id = ?", userID).
			First(&authenticatedUser)
	}

	err := db.Model(&article).
		Where("slug = ?", c.Params("slug")).
		Find(&article).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Redirect("/")
		}
	}

	return c.Render("articles/htmx-article-page", fiber.Map{
		"PageTitle":          article.Title,
		"NavBarActive":       "none",
		"Article":            article,
		"IsOob":              false,
		"IsSelf":             isSelf,
		"IsFollowed":         isFollowed,
		"IsArticleFavorited": false,
		"AuthenticatedUser":  authenticatedUser,
		"FiberCtx":           c,
	}, "layouts/app-htmx")
}

//#endregion
