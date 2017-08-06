package main

import (
	"github.com/goadesign/goa"
	"news-api/app"
)

// ArticlesController implements the articles resource.
type ArticlesController struct {
	*goa.Controller
}

// NewArticlesController creates a articles controller.
func NewArticlesController(service *goa.Service) *ArticlesController {
	return &ArticlesController{Controller: service.NewController("ArticlesController")}
}

// Article runs the article action.
func (c *ArticlesController) Article(ctx *app.ArticleArticlesContext) error {
	// ArticlesController_Article: start_implement

	// Put your logic here

	// ArticlesController_Article: end_implement
	res := &app.GoaNewsAPIArticle{}
	return ctx.OK(res)
}

// Articles runs the articles action.
func (c *ArticlesController) Articles(ctx *app.ArticlesArticlesContext) error {
	// ArticlesController_Articles: start_implement

	// Put your logic here

	// ArticlesController_Articles: end_implement
	res := app.GoaNewsAPIArticleCollection{}
	return ctx.OK(res)
}
