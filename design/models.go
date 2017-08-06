package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("news storage", func() {
	Description("news DBの操作")
	Store("MySQL", gorma.MySQL, func() {
		Model("Article", func() {
			// MediaTypeで生成したArticleにマッピング
			RendersTo(Article)
			Description("articles テーブル")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("head_line", gorma.String)
			Field("news_link", gorma.String)
			Field("thumb_url", gorma.String)
			Field("category_id", gorma.String)
			Field("media_id", gorma.String)
			Field("created", gorma.String)
			Field("updated", gorma.String)
			Field("vespa_index", gorma.String)
			Field("qsty_category_id", gorma.String)
			Field("description", gorma.String)
			Field("clicks", gorma.Integer)
			HasMany("Articles","Article")
		})
	})
})