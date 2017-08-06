// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "news api": Model Helpers
//
// Command:
// $ goagen
// --design=news-api/design
// --out=$(GOPATH)/src/news-api
// --version=v1.2.0-dirty

package models

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"news-api/app"
	"time"
)

// MediaType Retrieval Functions

// ListGoaNewsAPIArticle returns an array of view: default.
func (m *ArticleDB) ListGoaNewsAPIArticle(ctx context.Context) []*app.GoaNewsAPIArticle {
	defer goa.MeasureSince([]string{"goa", "db", "goaNewsAPIArticle", "listgoaNewsAPIArticle"}, time.Now())

	var native []*Article
	var objs []*app.GoaNewsAPIArticle
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Article", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.ArticleToGoaNewsAPIArticle())
	}

	return objs
}

// ArticleToGoaNewsAPIArticle loads a Article and builds the default view of media type GoaNewsApiArticle.
func (m *Article) ArticleToGoaNewsAPIArticle() *app.GoaNewsAPIArticle {
	article := &app.GoaNewsAPIArticle{}
	article.Clicks = m.Clicks
	article.Description = m.Description
	article.HeadLine = m.HeadLine
	article.ID = m.ID

	return article
}

// OneGoaNewsAPIArticle loads a Article and builds the default view of media type GoaNewsApiArticle.
func (m *ArticleDB) OneGoaNewsAPIArticle(ctx context.Context, id int) (*app.GoaNewsAPIArticle, error) {
	defer goa.MeasureSince([]string{"goa", "db", "goaNewsAPIArticle", "onegoaNewsAPIArticle"}, time.Now())

	var native Article
	err := m.Db.Scopes().Table(m.TableName()).Preload("Articles").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Article", "error", err.Error())
		return nil, err
	}

	view := *native.ArticleToGoaNewsAPIArticle()
	return &view, err
}
