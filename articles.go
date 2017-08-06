package main

import (
	"github.com/goadesign/goa"
	"news-api/app"
	"github.com/jinzhu/gorm"
	"news-api/models"
)

// ArticlesController implements the articles resource.
type ArticlesController struct {
	*goa.Controller
	db *gorm.DB
}

// NewArticlesController creates a articles controller.
func NewArticlesController(service *goa.Service, db *gorm.DB) *ArticlesController {
	return &ArticlesController{
		Controller: service.NewController("ArticlesController"),
		db:         db,
	}
}

// Article runs the article action.
func (c *ArticlesController) Article(ctx *app.ArticleArticlesContext) error {
	// ArticlesController_Article: start_implement

	// Put your logic here
	articleTable := models.NewArticleDB(c.db)
	article, err := articleTable.OneGoaNewsAPIArticle(ctx.Context, ctx.ID)

	if err != nil {
		return ctx.NotFound()
	}

	// ArticlesController_Article: end_implement
	res := &app.GoaNewsAPIArticle{}
	res = article
	return ctx.OK(res)
}

// Articles runs the articles action.
func (c *ArticlesController) Articles(ctx *app.ArticlesArticlesContext) error {
	// ArticlesController_Articles: start_implement

	// Put your logic here
	articleTable := models.NewArticleDB(c.db)
	articles := articleTable.ListGoaNewsAPIArticle(ctx.Context)

	// ArticlesController_Articles: end_implement
	res := app.GoaNewsAPIArticleCollection{}
	res = articles
	return ctx.OK(res)
}
