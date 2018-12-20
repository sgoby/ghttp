package ghttp

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
)

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Input          *RequestInput
	Params         httprouter.Params
	dataMap        map[string]interface{}
	nextServeHTTP  http.Handler
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	ctx := &Context{
		Request:        r,
		ResponseWriter: w,
		Input:          newInput(r),
		dataMap:        make(map[string]interface{}),
	}
	return ctx
}
//
func (ctx *Context) SetData(key string,val interface{}) {
	ctx.dataMap[key] = val
}
//
func (ctx *Context) GetData(key string) interface{}{
	return ctx.dataMap[key]
}
//
func (ctx *Context) SetParams(p httprouter.Params) {
	ctx.Params = p
}
//
func (ctx *Context) GetParams(name string) string {
	return ctx.Params.ByName(name)
}
//
func (ctx *Context) setNextServeHTTP(h http.Handler) {
	ctx.nextServeHTTP = h
}
//
func (ctx *Context) Next() {
	if ctx.nextServeHTTP != nil{
		ctx.nextServeHTTP.ServeHTTP(ctx.ResponseWriter,ctx.Request)
	}
}
//
func (ctx *Context) Redirect(path string) {
	http.Redirect(ctx.ResponseWriter, ctx.Request, path, http.StatusFound)
}

//
func (ctx *Context) Fprint(formart string, val ...interface{}) {
	ctx.ResponseWriter.Write([]byte(fmt.Sprintf(formart, val...)))
}
func (ctx *Context) Print(val ...interface{}) {
	ctx.ResponseWriter.Write([]byte(fmt.Sprint(val...)))
}

//
func (ctx *Context) Json(val interface{}) error {
	ctx.ResponseWriter.Header().Set("content-type", "application/json")
	paramJson, err := json.Marshal(val)
	if err != nil {
		return err
	}
	ctx.ResponseWriter.Write(paramJson)
	return nil
}
