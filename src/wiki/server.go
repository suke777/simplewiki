package wiki

import "net/http"

// Run ハンドラを登録してサーバーを起動します
func Run() {
	registerViewHandler()
	registerEditHandler()
	registerSaveHandler()
	registerErrorHandler()
	start()
}

func start() {
	http.ListenAndServe(":8080", nil)
}
