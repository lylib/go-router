package gorouter

import (
	"log"
	"net/http"
)

// 与原生http接口对接的
type handler struct {
	m map[string]http.HandlerFunc
}

func (this *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if f, ok := this.m[r.URL.Path]; ok {
		f.ServeHTTP(w, r)
		return
	}
	http.NotFound(w, r)
}

// 定制的业务路由
type router struct {
	h *handler
}

func NewRouter() *router {
	return &router{
		h: &handler{
			make(map[string]http.HandlerFunc),
		},
	}
}
func (this *router) Register(pattern string, handFunc http.HandlerFunc) {
	this.h.m[pattern] = handFunc
}
func (this *router) Run(ser *http.Server) {
	ser.Handler = this.h
	if err := ser.ListenAndServe(); nil != err {
		log.Fatal(err)
	}
}
