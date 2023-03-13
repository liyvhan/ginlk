package learn

import "net/http"

func http_net_01() {
	http.ListenAndServe(":80", nil)
}

/***
*http.ListenAndServe调用链
*1.创建Server结构体，调用Server.ListenAndServer
*2.for 监听连接
*3.对每个连接启动一个协程处理连接
*4.serverHandler{c.server}.ServeHTTP(w, w.req)处理请求，（自定义路由）
*5.如果没有设置handler，使用默认的DefaultServerMux(), key,val map处理请求
***/
