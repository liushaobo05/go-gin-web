package util

import (
	"reflect"
)

// 判断空
func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	// 优先通过断言来进行常用类型判断
	switch value := value.(type) {
	case int:
		return value == 0
	case int8:
		return value == 0
	case int16:
		return value == 0
	case int32:
		return value == 0
	case int64:
		return value == 0
	case uint:
		return value == 0
	case uint8:
		return value == 0
	case uint16:
		return value == 0
	case uint32:
		return value == 0
	case uint64:
		return value == 0
	case float32:
		return value == 0
	case float64:
		return value == 0
	case bool:
		return value == false
	case string:
		return value == ""
	case []byte:
		return len(value) == 0
	default:
		// Finally using reflect.
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Array:
			return rv.Len() == 0

		case reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return true
			}
		}
	}
	return false
}

//func Convert(i interface{}, t string) interface{} {
//	switch t {
//	case "int":
//		return Int(i)
//	case "int8":
//		return Int8(i)
//	case "int16":
//		return Int16(i)
//	case "int32":
//		return Int32(i)
//	case "int64":
//		return Int64(i)
//	case "uint":
//		return Uint(i)
//	case "uint8":
//		return Uint8(i)
//	case "uint16":
//		return Uint16(i)
//	case "uint32":
//		return Uint32(i)
//	case "uint64":
//		return Uint64(i)
//	case "float32":
//		return Float32(i)
//	case "float64":
//		return Float64(i)
//	case "bool":
//		return Bool(i)
//	case "string":
//		return String(i)
//	case "[]byte":
//		return Bytes(i)
//	case "[]int":
//		return Ints(i)
//	case "[]string":
//		return Strings(i)
//	default:
//		return i
//	}
//}

// to byte
//func Byte(i interface{}) byte {
//	if v, ok := i.(byte); ok {
//		return v
//	}
//	return byte(Uint8(i))
//}

// to rune
//func Rune(i interface{}) rune {
//	if v, ok := i.(rune); ok {
//		return v
//	}
//	return rune(Int32(i))
//}

//func Bytes(i interface{}) []byte {
//	if i == nil {
//		return nil
//	}
//	switch value := i.(type) {
//	case string:
//		return []byte(value)
//	case []byte:
//		return value
//	default:
//		return gbinary.Encode(i)
//	}
//}

// to []rune
//func Runes(i interface{}) []rune {
//	if v, ok := i.([]rune); ok {
//		return v
//	}
//	return []rune(String(i))
//}

// to string
//func String(i interface{}) string {
//	if i == nil {
//		return ""
//	}
//	switch value := i.(type) {
//	case int:
//		return strconv.FormatInt(int64(value), 10)
//	case int8:
//		return strconv.Itoa(int(value))
//	case int16:
//		return strconv.Itoa(int(value))
//	case int32:
//		return strconv.Itoa(int(value))
//	case int64:
//		return strconv.FormatInt(int64(value), 10)
//	case uint:
//		return strconv.FormatUint(uint64(value), 10)
//	case uint8:
//		return strconv.FormatUint(uint64(value), 10)
//	case uint16:
//		return strconv.FormatUint(uint64(value), 10)
//	case uint32:
//		return strconv.FormatUint(uint64(value), 10)
//	case uint64:
//		return strconv.FormatUint(uint64(value), 10)
//	case float32:
//		return strconv.FormatFloat(float64(value), 'f', -1, 32)
//	case float64:
//		return strconv.FormatFloat(value, 'f', -1, 64)
//	case bool:
//		return strconv.FormatBool(value)
//	case string:
//		return value
//	case []byte:
//		return string(value)
//	case []rune:
//		return string(value)
//	default:
//		if f, ok := value.(apiString); ok {
//			// If the variable implements the String() interface,
//			// then use that interface to perform the conversion
//			return f.String()
//		} else if f, ok := value.(apiError); ok {
//			// If the variable implements the Error() interface,
//			// then use that interface to perform the conversion
//			return f.Error()
//		} else {
//			// Finally we use json.Marshal to convert.
//			jsonContent, _ := json.Marshal(value)
//			return string(jsonContent)
//		}
//	}
//}

// Int converts <i> to int.
//func Int(i interface{}) int {
//	if i == nil {
//		return 0
//	}
//	if v, ok := i.(int); ok {
//		return v
//	}
//	return int(Int64(i))
//}

// Int8 converts <i> to int8.
//func Int8(i interface{}) int8 {
//	if i == nil {
//		return 0
//	}
//	if v, ok := i.(int8); ok {
//		return v
//	}
//	return int8(Int64(i))
//}

// Int16 converts <i> to int16.
//func Int16(i interface{}) int16 {
//	if i == nil {
//		return 0
//	}
//	if v, ok := i.(int16); ok {
//		return v
//	}
//	return int16(Int64(i))
//}

// Int32 converts <i> to int32.
//func Int32(i interface{}) int32 {
//	if i == nil {
//		return 0
//	}
//	if v, ok := i.(int32); ok {
//		return v
//	}
//	return int32(Int64(i))
//}

// Int64 converts <i> to int64.
//func Int64(i interface{}) int64 {
//	if i == nil {
//		return 0
//	}
//	switch value := i.(type) {
//	case int:
//		return int64(value)
//	case int8:
//		return int64(value)
//	case int16:
//		return int64(value)
//	case int32:
//		return int64(value)
//	case int64:
//		return value
//	case uint:
//		return int64(value)
//	case uint8:
//		return int64(value)
//	case uint16:
//		return int64(value)
//	case uint32:
//		return int64(value)
//	case uint64:
//		return int64(value)
//	case float32:
//		return int64(value)
//	case float64:
//		return int64(value)
//	case bool:
//		if value {
//			return 1
//		}
//		return 0
//	default:
//		s := String(value)
//		// Hexadecimal
//		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
//			if v, e := strconv.ParseInt(s[2:], 16, 64); e == nil {
//				return v
//			}
//		}
//		// Octal
//		if len(s) > 1 && s[0] == '0' {
//			if v, e := strconv.ParseInt(s[1:], 8, 64); e == nil {
//				return v
//			}
//		}
//		// Decimal
//		if v, e := strconv.ParseInt(s, 10, 64); e == nil {
//			return v
//		}
//		// Float64
//		return int64(Float64(value))
//	}
//}

// Uint converts <i> to uint.
//func Uint(i interface{}) uint {
//	if i == nil {
//		return 0
//	}
//	if v, ok := i.(uint); ok {
//		return v
//	}
//	return uint(Uint64(i))
//}
//
//// Uint8 converts <i> to uint8.
//func Uint8(i interface{}) uint8 {
//	if i == nil {
//		return 0
//	}
//	if v, ok := i.(uint8); ok {
//		return v
//	}
//	return uint8(Uint64(i))
//}
//
//// Uint16 converts <i> to uint16.
//func Uint16(i interface{}) uint16 {
//	if i == nil {
//		return 0
//	}
//	if v, ok := i.(uint16); ok {
//		return v
//	}
//	return uint16(Uint64(i))
//}
//
//// Uint32 converts <i> to uint32.
//func Uint32(i interface{}) uint32 {
//	if i == nil {
//		return 0
//	}
//	if v, ok := i.(uint32); ok {
//		return v
//	}
//	return uint32(Uint64(i))
//}
//
//// Uint64 converts <i> to uint64.
//func Uint64(i interface{}) uint64 {
//	if i == nil {
//		return 0
//	}
//	switch value := i.(type) {
//	case int:
//		return uint64(value)
//	case int8:
//		return uint64(value)
//	case int16:
//		return uint64(value)
//	case int32:
//		return uint64(value)
//	case int64:
//		return uint64(value)
//	case uint:
//		return uint64(value)
//	case uint8:
//		return uint64(value)
//	case uint16:
//		return uint64(value)
//	case uint32:
//		return uint64(value)
//	case uint64:
//		return value
//	case float32:
//		return uint64(value)
//	case float64:
//		return uint64(value)
//	case bool:
//		if value {
//			return 1
//		}
//		return 0
//	default:
//		s := String(value)
//		// Hexadecimal
//		if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
//			if v, e := strconv.ParseUint(s[2:], 16, 64); e == nil {
//				return v
//			}
//		}
//		// Octal
//		if len(s) > 1 && s[0] == '0' {
//			if v, e := strconv.ParseUint(s[1:], 8, 64); e == nil {
//				return v
//			}
//		}
//		// Decimal
//		if v, e := strconv.ParseUint(s, 10, 64); e == nil {
//			return v
//		}
//		// Float64
//		return uint64(Float64(value))
//	}
//}
//
//// Float32 converts <i> to float32.
//func Float32(i interface{}) float32 {
//	if i == nil {
//		return 0
//	}
//	if v, ok := i.(float32); ok {
//		return v
//	}
//	v, _ := strconv.ParseFloat(strings.TrimSpace(String(i)), 64)
//	return float32(v)
//}
//
//// Float64 converts <i> to float64.
//func Float64(i interface{}) float64 {
//	if i == nil {
//		return 0
//	}
//	if v, ok := i.(float64); ok {
//		return v
//	}
//	v, _ := strconv.ParseFloat(strings.TrimSpace(String(i)), 64)
//	return v
//}
