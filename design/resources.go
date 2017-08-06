package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// ものの定義
var _ = Resource("articles", func() {

	BasePath("/articles")

	Action("articles", func() {
		Description("記事の一覧")
		Routing(
			GET("/"),
		)
		Response(OK, Articles)
		Response(BadRequest, ErrorMedia)
	})

	Action("article", func() {
		Description("記事の詳細")
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id")
		})
		Response(OK, Article)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

// Swaggerをローカルで実行するめの定義
var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger/")
})
