package validator

import (
	"fmt"
	"regexp"
)

type IValidator interface {
	Verify(val string) error
}
//
var validatorMap map[string]IValidator
//
func init(){
	validatorMap = make(map[string]IValidator)
	validatorMap["required"] = Vrequired{}
	validatorMap["numeric"] = Vnumeric{}
	validatorMap["email"] = Vemail{}
	validatorMap["ip"] = Vip{}
	validatorMap["url"] = Vurl{}
	validatorMap["parameter"] = Vparameter{}
}
//
type Vrequired struct {}
type Vnumeric struct {}
type Vemail struct {}
type Vip struct {}
type Vurl struct {}
type Vparameter struct {}
//
func (v Vrequired)Verify(val string) error{
	if len(val) < 1{
		return fmt.Errorf("cannot be empty.")
	}
	return nil
}
//
func (v Vnumeric)Verify(val string) error{
	err := fmt.Errorf("must be a number.")
	if len(val) < 1{
		return err
	}
	ok,regerr := regexp.MatchString(`[^\D]`,val)
	if regerr != nil{
		return regerr
	}
	if !ok{
		return err
	}
	return nil
}
//
func (v Vemail)Verify(val string) error{
	err := fmt.Errorf("is not a valid mailbox.")
	if len(val) < 1{
		return err
	}
	ok,regerr := regexp.MatchString(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`,val)
	if regerr != nil{
		return regerr
	}
	if !ok{
		return err
	}
	return nil
}

//
func (v Vip)Verify(val string) error{
	err := fmt.Errorf("must be a valid IP address.")
	if len(val) < 1{
		return err
	}
	ok,regerr := regexp.MatchString(`^$(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`,val)
	if regerr != nil{
		return regerr
	}
	if !ok{
		return err
	}
	return nil
}

//
func (v Vurl)Verify(val string) error{
	err := fmt.Errorf("format is incorrect.")
	if len(val) < 1{
		return err
	}
	ok,regerr := regexp.MatchString(`[a-zA-z]+://[^\s]*`,val)
	if regerr != nil{
		return regerr
	}
	if !ok{
		return err
	}
	return nil
}

//
func (v Vparameter)Verify(val string) error{
	err := fmt.Errorf("format is incorrect.")
	if len(val) < 1{
		return err
	}
	ok,regerr := regexp.MatchString(`\W*$`,val)
	if regerr != nil{
		return regerr
	}
	if !ok{
		return err
	}
	return nil
}

//
func GetValidator(name string) IValidator{
	v,ok := validatorMap[name]
	if ok{
		return v
	}
	return nil
}



