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
	router.HandleFunc("/api/greet", greet)
	router.HandleFunc("/api/customers", getAllCustomers)
	router.HandleFunc("/api/customers_xml", getAllCustomers_xml)
	router.HandleFunc("/api/customers_flex", getAllCustomers_flex)
	router.HandleFunc("/api/customers/{customer_id}", getCustomer)
	router.HandleFunc("/api/customers_int/{customer_id:[0-9]+}", getCustomer_int)

	// HTTP動詞によるルーティング制御
	router.HandleFunc("/api/methods", ofGet).Methods(http.MethodGet)
	router.HandleFunc("/api/methods", ofPost).Methods(http.MethodPost)
	router.HandleFunc("/api/methods", ofPut).Methods(http.MethodPut)
	router.HandleFunc("/api/methods", ofDelete).Methods(http.MethodDelete)

	// 第一引数 -> リッスンするアドレス
	// 第二引数 -> 使用するmultiplexer | デフォルトで用意されているものを使用するため、nilを設定
	http.ListenAndServe("0.0.0.0:8080", router)
}
