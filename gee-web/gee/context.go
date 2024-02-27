package gee

import "net/http"

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	//request info
	Path   string
	Method string

	//response code
	StatusCode string
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}
