package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jeevic/7days-golang/gee-web/gee"
)

/*
(1)
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello Gee</h1>

(2)
$ curl "http://localhost:9999/hello?name=geektutu"
hello geektutu, you're at /hello

(3)
$ curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
{"password":"1234","username":"geektutu"}

(4)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/

func main() {
	r := gee.New()

	//r.GET("/index", func(c *gee.Context) {
	//	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	//})
	//v1 := r.Group("/v1")
	//{
	//	v1.GET("/", func(c *gee.Context) {
	//		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	//	})
	//
	//	v1.GET("/hello", func(c *gee.Context) {
	//		// expect /hello?name=geektutu
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//	})
	//}
	//v2 := r.Group("/v2")
	//{
	//	v2.GET("/hello/:name", func(c *gee.Context) {
	//		// expect /hello/geektutu
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//	})
	//	v2.POST("/login", func(c *gee.Context) {
	//		c.JSON(http.StatusOK, gee.H{
	//			"username": c.PostForm("username"),
	//			"password": c.PostForm("password"),
	//		})
	//	})
	//
	//}

	r.Use(gee.Logger()) // global midlleware
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}
	r.Run(":9999")
}

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
