package nautilus

import (
	"errors"
	"fmt"
	"reflect"
)

// FieldData containing struct field data such as name, type and tag
type FieldData struct {
	Name      string
	Type      reflect.Type
	Tags      reflect.StructTag
	Value     interface{}
	Exported  bool
	Anonymous bool
}

// GetType get the type name of a struct variable
// For example, if v is an int, GetType(v) return int as string
// Or if v is a struct or custom type, GetType(v) return type name
func GetType(v interface{}) string {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Struct {
		return t.Name()
	}

	panic("argument is not an struct")
}

// GetStructFieldsData analyze the given struct and return information
// about exported fields such as name, type and tag value
func GetStructFieldsData(i interface{}) ([]FieldData, error) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("argument is not a struct")
	}

	var ret []FieldData
	for i := 0; i < t.NumField(); i++ {
		fieldData := FieldData{}
		fieldData.Name = t.Field(i).Name
		fieldData.Type = t.Field(i).Type
		fieldData.Tags = t.Field(i).Tag
		fieldData.Exported = t.Field(i).PkgPath == ""
		fieldData.Anonymous = t.Field(i).Anonymous

		if fieldData.Exported {
			fieldData.Value = v.Field(i).Interface()
		} else {
			fieldData.Value = nil
		}

		ret = append(ret, fieldData)
	}

	return ret, nil
}

// PointerToValue converts given variable (if it is a pointer) to its value
func PointerToValue(i interface{}) interface{} {
	if reflect.ValueOf(i).Kind() == reflect.Ptr {
		t := reflect.TypeOf(i).Elem()

		// TODO: handle maps, arrays and slices
		switch t.Kind() {
		case reflect.Bool:
			ptr := i.(*bool)
			return *ptr
		case reflect.Int:
			ptr := i.(*int)
			return *ptr
		case reflect.Int8:
			ptr := i.(*int8)
			return *ptr
		case reflect.Int16:
			ptr := i.(*int16)
			return *ptr
		case reflect.Int32:
			ptr := i.(*int32)
			return *ptr
		case reflect.Int64:
			ptr := i.(*int64)
			return *ptr
		case reflect.Uint:
			ptr := i.(*uint)
			return *ptr
		case reflect.Uint8:
			ptr := i.(*uint8)
			return *ptr
		case reflect.Uint16:
			ptr := i.(*uint16)
			return *ptr
		case reflect.Uint32:
			ptr := i.(*uint32)
			return *ptr
		case reflect.Uint64:
			ptr := i.(*uint64)
			return *ptr
		case reflect.Float32:
			ptr := i.(*float32)
			return *ptr
		case reflect.Float64:
			ptr := i.(*float64)
			return *ptr
		case reflect.Complex64:
			ptr := i.(*complex64)
			return *ptr
		case reflect.Complex128:
			ptr := i.(*complex128)
			return *ptr
		case reflect.String:
			ptr := i.(*string)
			return *ptr
		case reflect.Struct:
			ptr := i.(*bool)
			return *ptr
		default:
			panic(fmt.Sprintf("Unsupported type [%s]", t.Kind().String()))
			return nil
		}
	}

	return i
}

// GetFieldPointer returns a pointer to given struct field with given field name
func GetFieldPointer(i interface{}, fieldName string) interface{} {
	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic("input must be a pointer to a struct")
	}

	fieldValue := v.Elem().FieldByName(fieldName)

	if fieldValue != (reflect.Value{}) {
		fieldInterface := fieldValue.Addr().Interface()
		return interfaceToTypePtr(fieldValue.Type(), fieldInterface)
	}

	return nil
}

func interfaceToTypePtr(t reflect.Type, i interface{}) interface{} {
	// TODO: handle maps, arrays and slices
	switch t.Kind() {
	case reflect.Bool:
		return i.(*bool)
	case reflect.Int:
		return i.(*int)
	case reflect.Int8:
		return i.(*int8)
	case reflect.Int16:
		return i.(*int16)
	case reflect.Int32:
		return i.(*int32)
	case reflect.Int64:
		return i.(*int64)
	case reflect.Uint:
		return i.(*uint)
	case reflect.Uint8:
		return i.(*uint8)
	case reflect.Uint16:
		return i.(*uint16)
	case reflect.Uint32:
		return i.(*uint32)
	case reflect.Uint64:
		return i.(*uint64)
	case reflect.Float32:
		return i.(*float32)
	case reflect.Float64:
		return i.(*float64)
	case reflect.Complex64:
		return i.(*complex64)
	case reflect.Complex128:
		return i.(*complex128)
	case reflect.String:
		return i.(*string)
	default:
		panic(fmt.Sprintf("Unsupported type [%s]", t.Kind().String()))
		return nil
	}
}

// SetFieldValue sets value of field with given name in given struct with given value
func SetFieldValue(i interface{}, field string, value interface{}) {
	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic("input must be a pointer to a struct")
	}

	fieldVal := v.Elem().FieldByName(field)

	switch fieldVal.Kind() {
	case reflect.Bool:
		fieldVal.SetBool(value.(bool))
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
		fieldVal.SetInt(value.(int64))
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
		fieldVal.SetUint(value.(uint64))
	case reflect.Float32:
	case reflect.Float64:
		fieldVal.SetFloat(value.(float64))
	case reflect.Complex64:
	case reflect.Complex128:
		fieldVal.SetComplex(value.(complex128))
	case reflect.String:
		fieldVal.SetString(value.(string))
	default:
		panic(fmt.Sprintf("Unsupported type [%s]", fieldVal.Kind().String()))
	}
}
