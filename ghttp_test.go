package ghttp

import (
	"testing"
	"fmt"
)

func Test_Server(t *testing.T){
	api := New()
	m,err := NewMiddleware("/api")
	if err != nil{
		fmt.Println(err)
		return
	}
	m2,err := NewMiddleware("/api2")
	if err != nil{
		fmt.Println(err)
		return
	}
	middlewareRouter1 := NewRouter()
	middlewareRouter1.POST("/test",test)
	middlewareRouter1.PUT("/test",test)
	middlewareRouter1.PATCH("/test",test)
	m.RegisterFunc(mid)
	m.LoadRouter(middlewareRouter1)
	//
	middlewareRouter2 := NewRouter()
	middlewareRouter2.GET("/test",test)
	m2.RegisterFunc(mid2)
	m2.LoadRouter(middlewareRouter2)
	//
	apiRouter := NewRouter()
	apiRouter.GET("/",index)
	//
	api.Use(m)
	api.Use(m2)
	api.LoadRouter(apiRouter)
	//
	api.ListenAndServe(":9595")
}

func index(ctx *Context){
	fmt.Fprint(ctx.ResponseWriter, "Welcome 0000!\n")
}
func mid(ctx *Context){
	id,_ := ctx.Input.Int("id")
	if id == 100{
		ctx.Next()
		return
	}
}
func mid2(ctx *Context){
	id,_ := ctx.Input.Validator("numeric").Int("id")
	if id == 200{
		ctx.Next()
		return
	}
	var pp = struct {
		Error string
	}{Error:"mid2 error"}
	err := ctx.Json(pp)
	if err != nil{
		fmt.Println(err)
	}
}
func test(ctx *Context){
	//id,_ := ctx.Input.Int("id")
	//fmt.Fprint(ctx.ResponseWriter, "test!\n")
	var pp = struct {
		Name string
	}{Name:"anjunsang"}
	err := ctx.Json(pp)
	if err != nil{
		fmt.Println(err)
	}
}
