package slice

import (
	"fmt"
	"reflect"
)

// StructPick pick a field value from struct slice, append to dest slice
// dest element's type should be the same as field value
// [{"name":"alice",age:12},{"name":"bob","age":22}] => ["alice", "bob"]
func StructPick(dest interface{}, src interface{}, fieldName string) error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr || reflect.TypeOf(dest).Elem().Kind() != reflect.Slice {
		return fmt.Errorf("dest should be ptr of slice")
	}
	if reflect.TypeOf(src).Kind() != reflect.Slice {
		return fmt.Errorf("src should be slice of struct")
	}
	sliceItemType := reflect.TypeOf(src).Elem()
	if sliceItemType.Kind() != reflect.Struct && sliceItemType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("src item should be struct or ptr of struct")
	}

	dv := reflect.ValueOf(dest).Elem()
	sv := reflect.ValueOf(src)
	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		field := v.FieldByName(fieldName)
		dv.Set(reflect.Append(dv, field))
	}
	return nil
}

// StructFill set field value from src slice
// [{"name":"",age:12},{"name":"","age":22}], ["alice", "bob"] => [{"name":"alice",age:12},{"name":"bob","age":22}]
func StructFill(dest interface{}, src interface{}, fieldName string) error {
	if reflect.TypeOf(dest).Kind() != reflect.Slice {
		return fmt.Errorf("dest should be slice of struct")
	}
	typ := reflect.TypeOf(dest).Elem()
	if typ.Kind() != reflect.Struct {
		return fmt.Errorf("slice item should be ptr of struct")
	}

	if reflect.TypeOf(src).Kind() != reflect.Slice {
		return fmt.Errorf("src should be slice")
	}

	iv := reflect.ValueOf(dest)
	vv := reflect.ValueOf(src)
	for i := 0; i < iv.Len(); i++ {
		v := iv.Index(i)
		field := v.FieldByName(fieldName)
		if field.IsValid() {
			field.Set(vv.Index(i))
		}
	}
	return nil
}
