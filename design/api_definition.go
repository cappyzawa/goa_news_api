package design                                     // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("news api", func() {
	// APIのタイトル
	Title("news api sample")
	// 説明
	Description("goaを用いたnews apiのサンプル")
	// コンタクト情報
	Contact(func() {
		Name("shu920921")
		Email("m153304w@st.u-gakugei.ac.jp")
		URL("https://github.com/shu920921/goa_news_api/issues")
	})
	// ライセンス
	License(func() {
		Name("MIT")
	})
	// ホストの設定
	Host("localhost:18080")
	//プロトコル定義
	Scheme("http", "https")
	// エンドポイントのベースパス
	BasePath("/api/v1")


})