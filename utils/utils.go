package utils

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	Computed = 0
	Optional = 1
	Required = 2
)

func setStatus(s *schema.Schema, status int) *schema.Schema {
	com := false
	opt := false
	req := false
	if status == Computed {
		com = true
	} else if status == Optional {
		opt = true
	} else {
		req = true
	}

	s.Computed = com
	s.Optional = opt
	s.Required = req

	return s
}

func RecurseSetSchemaStatus(r *schema.Resource, status int, parentSameStatus bool) *schema.Resource {
	for key, element := range r.Schema {
		isSetStatus := true
		if element.Type == schema.TypeList || element.Type == schema.TypeMap || element.Type == schema.TypeSet {
			if elemResource, ok := element.Elem.(*schema.Resource); ok {
				r.Schema[key].Elem = RecurseSetSchemaStatus(elemResource, status, parentSameStatus)
			}
			// } else if elemSchema, ok := element.Elem.(*schema.Schema); ok {
			// 	r.Schema[key].Elem = setStatus(elemSchema, status)
			// }
			isSetStatus = parentSameStatus
		}
		if isSetStatus {
			r.Schema[key] = setStatus(r.Schema[key], status)
		}
	}

	return r
}

func RecurseSetSchemaStatusByKey(r *schema.Resource, key string, status int, parentSameStatus bool) *schema.Resource {
	keys := strings.Split(key, ".")

	if _, ok := r.Schema[keys[0]]; !ok {
		fmt.Printf("[ERROR] key not found: %s\n", keys[0])
		return r
	}

	isSetStatus := true
	if len(keys) > 1 {
		r.Schema[keys[0]].Elem = RecurseSetSchemaStatusByKey(r.Schema[keys[0]].Elem.(*schema.Resource), strings.Join(keys[1:], "."), status, parentSameStatus)
		isSetStatus = parentSameStatus
	}

	if isSetStatus {
		r.Schema[keys[0]] = setStatus(r.Schema[keys[0]], status)
	}

	return r
}

func GenerateMapFromFields(t interface{}) map[string]interface{} {
	temp := make(map[string]interface{})

	for i := 0; i < reflect.TypeOf(t).NumField(); i++ {
		fieldName := reflect.TypeOf(t).Field(i).Tag.Get("json")
		if fieldName == "" {
			fieldName = reflect.TypeOf(t).Field(i).Name
		}
		fieldType := reflect.TypeOf(t).Field(i).Type.String()
		val := reflect.ValueOf(t).FieldByName(fieldName)

		if !val.IsValid() {
			log.Printf("[DEBUG] Val was invalid fieldName: %s fieldType: %s (%+v) with val %+v", fieldName, fieldType, fieldType, val)
			continue
		}

		if (len(fieldType) >= 5 && fieldType[0:6] == "client") || (len(fieldType) >= 7 && fieldType[0:8] == "[]client") {
			// log.Printf("[DEBUG] Got fieldName: %s fieldType: %s (%+v) with val %+v", fieldName, fieldType, fieldType, val)
			if val.Kind() == reflect.Struct {
				tempArray := make([]map[string]interface{}, 1)
				newVal := GenerateMapFromFields(val.Interface())
				if len(newVal) > 0 {
					tempArray[0] = newVal
					temp[strings.ToLower(fieldName)] = tempArray
				}
			} else if val.Kind() == reflect.Slice {
				tempArray := make([]map[string]interface{}, val.Len())
				atLeastOne := false
				for i := 0; i < val.Len(); i++ {
					newVal := GenerateMapFromFields(val.Index(i).Interface())
					if len(newVal) > 0 {
						atLeastOne = true
						tempArray[i] = newVal
					}
				}
				if atLeastOne {
					temp[strings.ToLower(fieldName)] = tempArray
				}
			} else {
				if val.IsNil() {
					continue
				}
				newVal := GenerateMapFromFields(val.Interface())
				if len(newVal) > 0 {
					temp[strings.ToLower(fieldName)] = newVal
				}
			}
		} else {
			switch fieldType {
			case
				"string":
				newVal := val.String()
				if len(newVal) > 0 {
					temp[strings.ToLower(fieldName)] = newVal
				}
			case "int":
				temp[strings.ToLower(fieldName)] = val.Int()
			case "bool":
				temp[strings.ToLower(fieldName)] = val.Bool()
			}
		}
	}

	// log.Printf("[DEBUG] Temp: %+v", temp)

	return temp
}
