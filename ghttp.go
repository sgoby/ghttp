package ghttp

import (
	"net/http"
	"log"
	"fmt"
	"time"
)

type GHttp struct {
	*Router
	middlewares Middlewares
	gLogger     *log.Logger
}

type ResponseWriter struct {
	responseWriter http.ResponseWriter
	statusCode     int
	contentLength  int
	ctx            *Context
}

//
func (rw *ResponseWriter) setContext(ctx *Context) {
	rw.ctx = ctx
}

//
func (rw *ResponseWriter) getContext() (ctx *Context) {
	return rw.ctx
}

//
func (rw *ResponseWriter) Header() http.Header {
	return rw.responseWriter.Header()
}

//Write([]byte) (int, error)
func (rw *ResponseWriter) Write(val []byte) (len int, err error) {
	len, err = rw.responseWriter.Write(val)
	if err != nil {
		return
	}
	rw.contentLength += len
	return
}

//WriteHeader(int)
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.responseWriter.WriteHeader(statusCode)
}

type Method func(ctx *Context)

/*
 log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                   '"$status" $body_bytes_sent "$http_referer" '
                   '"$http_user_agent" "$http_x_forwarded_for" '
                   '"$gzip_ratio" $request_time $bytes_sent $request_length';
*/
var accessFormat string = "%s - - [%s] \"%s %s %s\" - %v %v - %s %.2fs"
//
func New() *GHttp {
	g := &GHttp{}
	g.Router = NewRouter()
	return g
}

//
func (g *GHttp) Use(m *Middleware) {
	g.middlewares = append(g.middlewares, m)
}

//
func (g *GHttp) LoadRouter(r *Router) {
	g.Router = r
}

//
func (g *GHttp) SetAccessLogger(plog *log.Logger) {
	g.gLogger = plog
}
//
func (g *GHttp)ListenAndVerifyClientCert(addr,caFile,certFile, keyFile string) error{
	for _, m := range g.middlewares {
		m.initRouter()
	}
	g.initRouter()
	httpServer,err := g.getHttpServerTLS(addr,caFile)
	if err != nil{
		return err
	}
	//
	return httpServer.ListenAndServeTLS(certFile, keyFile)
}
//https
func (g *GHttp)ListenAndServeTLS(addr,certFile, keyFile string) error{
	for _, m := range g.middlewares {
		m.initRouter()
	}
	g.initRouter()
	httpServer,err := g.getHttpServer(addr)
	if err != nil{
		return err
	}
	//
	return httpServer.ListenAndServeTLS(certFile, keyFile)
}
//
func (g *GHttp) ListenAndServe(addr string) {
	for _, m := range g.middlewares {
		m.initRouter()
	}
	g.initRouter()
	http.ListenAndServe(addr, g)
}
//
func (g *GHttp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	respW := &ResponseWriter{
		responseWriter: w,
		statusCode:     http.StatusOK,
	}
	beginTime := time.Now()
	defer func() {
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		userTime := time.Now().UnixNano() - beginTime.UnixNano()
		ms := float64(userTime) / float64(time.Second)
		accessLog := fmt.Sprintf(accessFormat, req.RemoteAddr, nowTime, req.Method, req.RequestURI, req.Proto, respW.statusCode, respW.contentLength, req.UserAgent(), ms)
		if g.gLogger != nil {
			g.gLogger.Println(accessLog)
		}
	}()
	//
	middlewareCount := 0
	for _, m := range g.middlewares {
		if m.matchPath(req.URL.Path) {
			m.ServeHTTP(respW, req)
			middlewareCount += 1
			if m.len() > 0 {
				break
			}
			continue
		}
	}
	if middlewareCount <= 0 {
		g.Router.ServeHTTP(respW, req)
	}
}


