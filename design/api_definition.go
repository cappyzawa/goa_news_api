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
		URL("https://github.com/shu920921/goa-news-api/issues")
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

	// CORSポリシー
	Origin("http://localhost:18080/swagger", func() {
		//ヘッダー
		Expose("X-Time")
		//許可するHTTPメソッド
		//現状GETしか使わないが。。。
		Methods("GET", "POST", "PUT", "DELETE")
		//キャッシュする時間
		MaxAge(600)
		// Access-Control-Allow-Credentialsヘッダーを設定する
		Credentials()
	})
})