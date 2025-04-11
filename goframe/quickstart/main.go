package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
    s.BindHandler("/", func(r *ghttp.Request) {
        r.Response.Writef(
            "Hello %s! Your Age is %d",
            r.Get("name", "unknown").String(),
            r.Get("age").Int(),
        )
    })
    s.SetPort(8000)
    s.Run()
}