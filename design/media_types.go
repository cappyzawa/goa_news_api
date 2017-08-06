package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//レスポンスデータの定義
var Article = MediaType("application/vnd.goa.news.api.article+json", func() {
	Description("記事詳細のデータ")
	Attributes(func() {
		Attribute("id", Integer)
		Attribute("head_line", String)
		Attribute("news_link", String)
		Attribute("thumb_url", String)
		Attribute("category_id", String)
		Attribute("media_id", String)
		Attribute("created", String)
		Attribute("updated", String)
		Attribute("vespa_index", String)
		Attribute("qsty_category_id", String)
		Attribute("description", String)
		Attribute("clicks", Integer)
		Required(
			"id",
			"head_line",
			"news_link",
			"thumb_url",
			"category_id",
			"media_id",
			"created",
			"updated",
			"vespa_index",
			"qsty_category_id",
			"description",
			"clicks",
		)
	})
	View("default", func() {
		Attribute("id")
		Attribute("head_line")
		Attribute("description")
		Attribute("clicks")
	})
})

var Articles = MediaType("application/vnd.goa.news.api.articles+json", func() {
	Description("記事一覧のデータ")
	Attribute("articles", ArrayOf(Article))
	View("default", func() {
		Attribute("articles")
	})
})