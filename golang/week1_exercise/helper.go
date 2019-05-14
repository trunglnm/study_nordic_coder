package exercise

import (
	"reflect"
	"strings"
)

/**
* clone Lodash check object is empty
* input: obj [numberic // boolean // array // struct]
* output: bool [true/false]
**/
func IsEmpty(obj interface{}) bool {
	if obj == nil || obj == false || obj == "" || obj == 0 {
		return true
	}
	valueOfObj := reflect.ValueOf(obj)
	if valueOfObj.Kind() == reflect.Array || valueOfObj.Kind() == reflect.Slice {
		if valueOfObj.Len() == 0 {
			return true
		}
	}
	if valueOfObj.Kind() == reflect.Ptr {
		valueOfObj = valueOfObj.Elem()
	}
	if valueOfObj.Kind() == reflect.Struct {
		for i := 0; i < valueOfObj.NumField(); i++ {
			field := valueOfObj.Field(i)
			valueOfField := field.Interface()
			if !(valueOfField == nil || valueOfField == false || valueOfField == "" || valueOfField == 0) {
				return false
			}
		}
		return true
	}
	return false
}

/**
* clone Lodash find maxinum
* input: obj [array]
* output: bool [true/false]
**/
func Max(arr interface{}) interface{} {
	valueOfObj := reflect.ValueOf(arr)
	typeOfObj := reflect.TypeOf(arr)
	if valueOfObj.Kind() != reflect.Array && valueOfObj.Kind() != reflect.Slice {
		panic("make sure input is array or slice")
	}
	if valueOfObj.Len() == 0 {
		panic("make sure input is not empty array")
	}

	max := valueOfObj.Index(0)
	for i := 1; i < valueOfObj.Len(); i++ {
		switch typeOfObj.Elem().Kind() {
		case reflect.Int:
			if max.Int() < valueOfObj.Index(i).Int() {
				max = valueOfObj.Index(i)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if max.Uint() < valueOfObj.Index(i).Uint() {
				max = valueOfObj.Index(i)
			}
		case reflect.Float32, reflect.Float64:
			if max.Float() < valueOfObj.Index(i).Float() {
				max = valueOfObj.Index(i)
			}
		}
	}

	return max.Interface()
}

/**
* clone Lodash check contain
* input: obj [array / string / struct]
* output: bool [true/false]
**/
func Contain(obj interface{}, element interface{}) bool {
	typeOfObj := reflect.TypeOf(obj)
	valueOfObj := reflect.ValueOf(obj)
	typeOfElement := reflect.TypeOf(element)
	switch typeOfObj.Kind() {
	case reflect.Array, reflect.Slice:
		if valueOfObj.Len() == 0 {
			return false
		}
		if typeOfElement.Kind() == reflect.Array || typeOfElement.Kind() == reflect.Slice {
			return CheckArrContainNested(obj, element)
		}
		return CheckArrContainSimple(obj, element)
	case reflect.String:
		if typeOfElement.Kind() != reflect.String {
			panic("make sure element is same type with Obj | String")
		}
		valueOfObj := reflect.ValueOf(obj).String()
		valueOfElement := reflect.ValueOf(element).String()
		return strings.Index(valueOfObj, valueOfElement) != -1
	case reflect.Struct:
		for i := 0; i < valueOfObj.NumField(); i++ {
			field := valueOfObj.Field(i)
			valueOfField := field.Interface()
			if typeOfElement.Kind() != field.Type().Kind() {
				continue
			}
			switch typeOfElement.Kind() {
			case reflect.Int:
				if reflect.ValueOf(valueOfField).Int() == reflect.ValueOf(element).Int() {
					return true
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if reflect.ValueOf(valueOfField).Uint() == reflect.ValueOf(element).Uint() {
					return true
				}
			case reflect.Float32, reflect.Float64:
				if reflect.ValueOf(valueOfField).Float() == reflect.ValueOf(element).Float() {
					return true
				}
			case reflect.String:
				if reflect.ValueOf(valueOfField).String() == reflect.ValueOf(element).String() {
					return true
				}
			}
		}
	default:
		panic("make sure obj is array or string or struct")
	}

	return false
}

func CheckArrContainSimple(arr interface{}, element interface{}) bool {
	typeOfArr := reflect.TypeOf(arr)
	typeOfElement := reflect.TypeOf(element)
	if typeOfElement.Kind() != typeOfArr.Elem().Kind() {
		panic("make sure element is same type with Array")
	}

	valueOfArr := reflect.ValueOf(arr)
	valueOfElement := reflect.ValueOf(element)
	for i := 0; i < valueOfArr.Len(); i++ {
		switch typeOfElement.Kind() {
		case reflect.Int:
			if valueOfElement.Int() == valueOfArr.Index(i).Int() {
				return true
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if valueOfElement.Uint() == valueOfArr.Index(i).Uint() {
				return true
			}
		case reflect.Float32, reflect.Float64:
			if valueOfElement.Float() == valueOfArr.Index(i).Float() {
				return true
			}
		case reflect.String:
			if valueOfElement.String() == valueOfArr.Index(i).String() {
				return true
			}
		}
	}
	return false
}

func CheckArrContainNested(arr interface{}, element interface{}) bool {
	valueOfElement := reflect.ValueOf(element)
	if valueOfElement.Len() == 0 {
		return false
	}
	for i := 0; i < valueOfElement.Len(); i++ {
		if !CheckArrContainSimple(arr, valueOfElement.Index(i).Interface()) {
			return false
		}
	}
	return true
}

/**
* clone Lodash get last element
* input: obj [array]
* output: element
**/
func Last(obj interface{}) interface{} {
	typeOfObj := reflect.TypeOf(obj)
	valueOfObj := reflect.ValueOf(obj)
	if typeOfObj.Kind() != reflect.Array && typeOfObj.Kind() != reflect.Slice {
		panic("make sure obj is array or slice")
	}
	if valueOfObj.Len() == 0 {
		panic("make sure obj is array not empty")
	}

	return valueOfObj.Index(valueOfObj.Len() - 1).Interface()
}
