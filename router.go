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
	routerMap        map[string][]*routerStmt // method
}

//
type routerStmt struct {
	urlPath      string
	handleMethod Method
}

//
func (rs *routerStmt) handle(w http.ResponseWriter,r *http.Request,argus httprouter.Params){
	var ctx *Context
	if customizeW,ok := w.(*ResponseWriter);ok{
		ctx = customizeW.getContext()
	}
	if ctx == nil{
		ctx = NewContext(w, r)
	}
	ctx.SetParams(argus)
	rs.handleMethod(ctx)
}
//
func NewRouter() *Router {
	return &Router{
		routerMap: make(map[string][]*routerStmt),
	}
}

//
func (r *Router) len() int {
	return len(r.routerMap)
}

//
func (r *Router) setOptimizationPath(optiFunc func(string) string) {
	r.optimizationPath = optiFunc
}

//
func (r *Router) initRouter() error {
	for k, rStmts := range r.routerMap {
		switch k {
		case strGET:
			for _, rStmt := range rStmts {
				if r.optimizationPath != nil{
					rStmt.urlPath = r.optimizationPath(rStmt.urlPath)
				}
				r.Router.GET(rStmt.urlPath,rStmt.handle)
			}
		case strPOST:
			for _, rStmt := range rStmts {
				if r.optimizationPath != nil{
					rStmt.urlPath = r.optimizationPath(rStmt.urlPath)
				}
				r.Router.POST(rStmt.urlPath,rStmt.handle)
			}
		case strHEAD:
			for _, rStmt := range rStmts {
				if r.optimizationPath != nil{
					rStmt.urlPath = r.optimizationPath(rStmt.urlPath)
				}
				r.Router.HEAD(rStmt.urlPath,rStmt.handle)
			}
		case strOPTIONS:
			for _, rStmt := range rStmts {
				if r.optimizationPath != nil{
					rStmt.urlPath = r.optimizationPath(rStmt.urlPath)
				}
				r.Router.OPTIONS(rStmt.urlPath,rStmt.handle)
			}
		case strPUT:
			for _, rStmt := range rStmts {
				if r.optimizationPath != nil{
					rStmt.urlPath = r.optimizationPath(rStmt.urlPath)
				}
				r.Router.PUT(rStmt.urlPath,rStmt.handle)
			}
		case strPATCH:
			for _, rStmt := range rStmts {
				if r.optimizationPath != nil{
					rStmt.urlPath = r.optimizationPath(rStmt.urlPath)
				}
				r.Router.PATCH(rStmt.urlPath,rStmt.handle)
			}
		case strDELETE:
			for _, rStmt := range rStmts {
				if r.optimizationPath != nil{
					rStmt.urlPath = r.optimizationPath(rStmt.urlPath)
				}
				r.Router.DELETE(rStmt.urlPath,rStmt.handle)
			}
		default:
			return fmt.Errorf("Invalid request method: %v", k)
		}
	}
	return nil
}

//
func (r *Router) GET(path string, m Method) {
	r.routerMap[strGET] = append(r.routerMap[strGET], &routerStmt{
		urlPath:      path,
		handleMethod: m,
	})
}
func (r *Router) POST(path string, m Method) {
	r.routerMap[strPOST] = append(r.routerMap[strPOST], &routerStmt{
		urlPath:      path,
		handleMethod: m,
	})
}
func (r *Router) HEAD(path string, m Method) {
	r.routerMap[strHEAD] = append(r.routerMap[strHEAD], &routerStmt{
		urlPath:      path,
		handleMethod: m,
	})
}
func (r *Router) OPTIONS(path string, m Method) {
	r.routerMap[strOPTIONS] = append(r.routerMap[strOPTIONS], &routerStmt{
		urlPath:      path,
		handleMethod: m,
	})
}
func (r *Router) PUT(path string, m Method) {
	r.routerMap[strPUT] = append(r.routerMap[strPUT], &routerStmt{
		urlPath:      path,
		handleMethod: m,
	})
}
func (r *Router) PATCH(path string, m Method) {
	r.routerMap[strPATCH] = append(r.routerMap[strPATCH], &routerStmt{
		urlPath:      path,
		handleMethod: m,
	})
}
func (r *Router) DELETE(path string, m Method) {
	r.routerMap[strDELETE] = append(r.routerMap[strDELETE], &routerStmt{
		urlPath:      path,
		handleMethod: m,
	})
}
