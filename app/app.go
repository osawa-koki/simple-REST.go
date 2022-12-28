package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// 自身でmultiplexerを登録
	router := mux.NewRouter()

	// ハンドラファンクション -> デフォルトmultiplexerに登録。 | 自身で作成したmultiplexerを使用
	// 第一引数 -> パターン
	// 第二引数 -> 「レスポンスライター」「リクエスト」を引数として受け取る関数
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)
	router.HandleFunc("/customers_xml", getAllCustomers_xml)
	router.HandleFunc("/customers_flex", getAllCustomers_flex)
	router.HandleFunc("/customers/{customer_id}", getCustomer)
	router.HandleFunc("/customers_int/{customer_id:[0-9]+}", getCustomer_int)

	// HTTP動詞によるルーティング制御
	router.HandleFunc("/methods", ofGet).Methods(http.MethodGet)
	router.HandleFunc("/methods", ofPost).Methods(http.MethodPost)
	router.HandleFunc("/methods", ofPut).Methods(http.MethodPut)
	router.HandleFunc("/methods", ofDelete).Methods(http.MethodDelete)

	// 第一引数 -> リッスンするアドレス
	// 第二引数 -> 使用するmultiplexer | デフォルトで用意されているものを使用するため、nilを設定
	http.ListenAndServe("0.0.0.0:80", router)
}
