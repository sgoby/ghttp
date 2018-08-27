package ghttp

import (
	"net/http"
	"errors"
)

//
type MiddlewareMethod func(ctx *Context)
//
type Middlewares []*Middleware
func(ms Middlewares)Len() int{
	return len(ms)
}
func(ms Middlewares)Less(i, j int) bool{
	if ms[i].currentPath < ms[j].currentPath{
		return false
	}
	return true
}
func(ms Middlewares)Swap(i, j int){
	ms[i],ms[j] = ms[j],ms[i]
}

type Middleware struct {
	*Router
	currentPath string
	middFunc MiddlewareMethod
}
//
func NewMiddleware(path string) (*Middleware,error){
	if path[0] != '/' {
		return nil,errors.New("path must begin with '/' in path '" + path + "'")
	}
	if path[len(path) -1] != '/'{
		path += "/"
	}
	m := &Middleware{
		currentPath:path,
	}
	//
	m.Router = NewRouter()
	m.Router.setOptimizationPath(m.optimizationPath)
	return m,nil
}
//
func (m *Middleware) ServeHTTP(w *ResponseWriter, req *http.Request){
	ctx := NewContext(w,req)
	ctx.setNextServeHTTP(m.Router)
	m.middFunc(&ctx)
}
//
func(m *Middleware) matchPath(path string) bool{
	if len(path) > len(m.currentPath)  && path[0:len(m.currentPath)] == m.currentPath{
		return true;
	}
	return false;
}
//
func(m *Middleware) getCurrentPath(path string) string{
	return m.currentPath
}
func(m *Middleware) setCurrentPath(path string){
	m.currentPath = path
}
//
func(m *Middleware) RegisterFunc(f MiddlewareMethod){
	m.middFunc = f
}
//
func(m *Middleware) LoadRouter(r *Router){
	m.Router = r
	m.Router.setOptimizationPath(m.optimizationPath)
}
//
func (m *Middleware) optimizationPath(path string) string{
	if path[0] == '/' {
		path = path[1:]
	}
	return m.currentPath + path
}

