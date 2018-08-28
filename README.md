# Restful Api 开发框架

Golang 轻量级Restful Api开发框架。支持中间件。自带参数验证器(Validator), 可自由扩展。

### 使用

```go
import "github.com/sgoby/ghttp"

...

func Test_Server(t *testing.T){
	api := ghttp.New()
	m,err := ghttp.NewMiddleware("/api")
	if err != nil{
		fmt.Println(err)
		return
	}
	middlewareRouter1 := ghttp.NewRouter()
	middlewareRouter1.POST("/test",test)
	//
	m.RegisterFunc(middlewareFunc)
	m.LoadRouter(middlewareRouter1)
	//
	apiRouter := ghttp.NewRouter()
	apiRouter.GET("/",index)
	//use middleware
	api.Use(m)
	api.LoadRouter(apiRouter)
	//
	api.ListenAndServe(":9595")
}
//
func index(ctx *ghttp.Context){
	fmt.Fprint(ctx.ResponseWriter, "Welcome!\n")
}
//Middleware handle function
func middlewareFunc(ctx *ghttp.Context){
	id,_ := ctx.Input.Validator("numeric").Int("id")
	if id == 200{
		ctx.Next()
		return
	}
	var pp = struct {
		Error string
	}{Error:"mid2 error"}
	ctx.Json(pp)
}
```