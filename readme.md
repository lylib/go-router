```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	gorouter "go-router"
)

func main() {
	r := gorouter.NewRouter()
	r.Register("/hello", func(w http.ResponseWriter, r *http.Request) {
		// 读取
		fmt.Println("=============================")
		// header携带的部分有用信息
		fmt.Println("RequestURI", r.RequestURI)
		fmt.Println("RemoteAddr", r.RemoteAddr)
		fmt.Println("UserAgent", r.UserAgent())
		fmt.Println("Referer", r.Referer())

		// 无论Content-Type都会传入并读取
		fmt.Println("URL", r.URL.Query().Get("G1")) // URL里面的参数

		// Content-Type:  application/...
		// [默认会去读URL参数(会覆盖form-data同字段的数据)] [x-www-form-urlencoded]
		fmt.Println("FormValue", r.FormValue("G1"))
		// [form-data] [x-www-form-urlencoded]
		fmt.Println("PostFormValue", r.PostFormValue("G1"))
		// [multipart/form-data]
		r.MultipartReader() // 读文件用的
		// [application/json]
		fmt.Println(ioutil.ReadAll(r.Body)) //

		// 返回
		// w.Header().Add("", "")
		w.Write([]byte(`{"B1":"Y1"}`))
	})
	r.Run(&http.Server{
		Addr:           ":8086",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})
}
```