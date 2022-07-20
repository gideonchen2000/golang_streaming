package main

import (
	"golang_streaming/video_server/streamserver/handlers"
	"golang_streaming/video_server/streamserver/limiter"
	"golang_streaming/video_server/streamserver/response"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *limiter.ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = limiter.NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		response.SendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)

	defer m.l.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", handlers.StreamHandler)

	router.POST("/upload/:vid-id", handlers.UploadHandler)

	router.GET("/testpage", handlers.TestPageHandler)

	return router
}
func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 100) // 流控值暂置为2便于测试
	http.ListenAndServe(":9000", mh)
}
