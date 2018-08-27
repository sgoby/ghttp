package ghttp

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

const (
	strGET     = "GET"
	strPOST    = "POST"
	strHEAD    = "HEAD"
	strOPTIONS = "OPTIONS"
	strPUT     = "PUT"
	strPATCH   = "PATCH"
	strDELETE  = "DELETE"
)

//
type Router struct {
	httprouter.Router
	optimizationPath func(string) string
	routerMap        map[string]*routerStmt // method
}

//
type routerStmt struct {
	urlPath      string
	handleMethod Method
}

//
func NewRouter() *Router {
	return &Router{
		routerMap:make(map[string]*routerStmt),
	}
}
//
func (r *Router) len() int{
	return len(r.routerMap)
}
//
func (r *Router) setOptimizationPath(optiFunc func(string) string) {
	r.optimizationPath = optiFunc
}

//
func (r *Router) initRouter() error{
	for k,rStmt := range r.routerMap{
		if r.optimizationPath != nil {
			rStmt.urlPath = r.optimizationPath(rStmt.urlPath)
		}
		//
		switch k {
		case strGET:
			r.Router.GET(rStmt.urlPath, func(w http.ResponseWriter, r *http.Request, argus httprouter.Params) {
				ctx := NewContext(w, r)
				ctx.SetParams(argus)
				rStmt.handleMethod(&ctx)
			})
		case strPOST:
			r.Router.POST(rStmt.urlPath, func(w http.ResponseWriter, r *http.Request, argus httprouter.Params) {
				ctx := NewContext(w, r)
				ctx.SetParams(argus)
				rStmt.handleMethod(&ctx)
			})
		case strHEAD:
			r.Router.HEAD(rStmt.urlPath, func(w http.ResponseWriter, r *http.Request, argus httprouter.Params) {
				ctx := NewContext(w, r)
				ctx.SetParams(argus)
				rStmt.handleMethod(&ctx)
			})
		case strOPTIONS:
			r.Router.OPTIONS(rStmt.urlPath, func(w http.ResponseWriter, r *http.Request, argus httprouter.Params) {
				ctx := NewContext(w, r)
				ctx.SetParams(argus)
				rStmt.handleMethod(&ctx)
			})
		case strPUT:
			r.Router.PUT(rStmt.urlPath, func(w http.ResponseWriter, r *http.Request, argus httprouter.Params) {
				ctx := NewContext(w, r)
				ctx.SetParams(argus)
				rStmt.handleMethod(&ctx)
			})
		case strPATCH:
			r.Router.PATCH(rStmt.urlPath, func(w http.ResponseWriter, r *http.Request, argus httprouter.Params) {
				ctx := NewContext(w, r)
				ctx.SetParams(argus)
				rStmt.handleMethod(&ctx)
			})
		case strDELETE:
			r.Router.DELETE(rStmt.urlPath, func(w http.ResponseWriter, r *http.Request, argus httprouter.Params) {
				ctx := NewContext(w, r)
				ctx.SetParams(argus)
				rStmt.handleMethod(&ctx)
			})
		default:
			return fmt.Errorf("Invalid request method: %v",k)
		}
	}
	return nil
}
//
func (r *Router) GET(path string, m Method) {
	r.routerMap[strGET] = &routerStmt{
		urlPath:path,
		handleMethod:m,
	}
}
func (r *Router) POST(path string, m Method) {
	r.routerMap[strPOST] = &routerStmt{
		urlPath:path,
		handleMethod:m,
	}
}
func (r *Router) HEAD(path string, m Method) {
	r.routerMap[strHEAD] = &routerStmt{
		urlPath:path,
		handleMethod:m,
	}
}
func (r *Router) OPTIONS(path string, m Method) {
	r.routerMap[strOPTIONS] = &routerStmt{
		urlPath:path,
		handleMethod:m,
	}
}
func (r *Router) PUT(path string, m Method) {
	r.routerMap[strPUT] = &routerStmt{
		urlPath:path,
		handleMethod:m,
	}
}
func (r *Router) PATCH(path string, m Method) {
	r.routerMap[strPATCH] = &routerStmt{
		urlPath:path,
		handleMethod:m,
	}
}
func (r *Router) DELETE(path string, m Method) {
	r.routerMap[strDELETE] = &routerStmt{
		urlPath:path,
		handleMethod:m,
	}
}
