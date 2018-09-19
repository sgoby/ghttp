package ghttp

import (
	"net/http"
	"strconv"
	"github.com/sgoby/ghttp/validator"
	"fmt"
	"net/url"
)
//
type RequestInput struct {
	request   *http.Request
	validator  validator.IValidator
}

//
func newInput(r *http.Request) *RequestInput {
	input := &RequestInput{
		request: r,
	}
	return input;
}

//rule [email]
func (i RequestInput) Validator(v interface{}) (r *RequestInput) {
	if iv,ok := v.(validator.IValidator);ok{
		i.validator = iv
	}
	if name,ok := v.(string);ok{
		i.validator = validator.GetValidator(name)
	}
	return &i
}
//
func (i *RequestInput) String(key string, defaultVals ...string) (string, error) {
	val := i.get(key);
	if i.validator != nil{
		if err := i.validator.Verify(val);err != nil{
			return "",fmt.Errorf("%s %v",key,err)
		}
	}
	if len(val) < 1 && len(defaultVals) > 0 {
		return defaultVals[0],nil
	}
	return val,nil
}
func (i *RequestInput) Int(key string, defaultVals ...int) (int, error) {
	val := i.get(key);
	if i.validator != nil{
		if err := i.validator.Verify(val);err != nil{
			return 0,fmt.Errorf("%s %v",key,err)
		}
	}
	if len(val) < 1 {
		if len(defaultVals) > 0 {
			return defaultVals[0], nil
		}
		return 0, nil
	}
	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err;
	}
	return int(num), nil
}
func (i *RequestInput) Int64(key string, defaultVals ...int64) (int64, error) {
	val := i.get(key);
	if i.validator != nil{
		if err := i.validator.Verify(val);err != nil{
			return 0,fmt.Errorf("%s %v",key,err)
		}
	}
	if len(val) < 1 {
		if len(defaultVals) > 0 {
			return defaultVals[0], nil
		}
		return 0, nil
	}
	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err;
	}
	return num, nil
}
func (i *RequestInput) Float32(key string, defaultVals ...float32) (float32, error) {
	val := i.get(key);
	if len(val) < 1 {
		if len(defaultVals) > 0 {
			return defaultVals[0], nil
		}
		return 0, nil
	}
	num, err := strconv.ParseFloat(val, 64)
	return float32(num), err
}
func (i *RequestInput) Float64(key string, defaultVals ...float64) (float64, error) {
	val := i.get(key);
	if len(val) < 1 {
		if len(defaultVals) > 0 {
			return defaultVals[0], nil
		}
		return 0, nil
	}
	return strconv.ParseFloat(val, 64)
}
func (i *RequestInput) Bool(key string, defaultVals ...bool) (bool, error) {
	val := i.get(key);
	if len(val) < 1 {
		if len(defaultVals) > 0 {
			return defaultVals[0], nil
		}
		return false, nil
	}
	return strconv.ParseBool(val)
}
func (i *RequestInput) Slice(key string, defaultVals ...bool) []interface{} {
	return nil
}

//
func (i *RequestInput) Form() url.Values {
	if i.request.Form == nil {
		i.request.ParseMultipartForm(32 << 20)
	}
	return i.request.Form
}

//
func (i *RequestInput) FormValue(key string, defaultVals ...string) string {
	val := i.request.FormValue(key)
	if len(val) < 1 && len(defaultVals) > 0 {
		return defaultVals[0]
	}
	return val
}

//
func (i *RequestInput) get(key string) string {
	if i.request == nil {
		return ""
	}
	if i.request.URL != nil {
		values := i.request.URL.Query()
		val := values.Get(key)
		if len(val) > 0 {
			return val
		}
	}
	val := i.request.FormValue(key)
	return val;
}
