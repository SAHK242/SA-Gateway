package grpcutil

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	TypeString  = reflect.TypeOf("")
	TypeBool    = reflect.TypeOf(true)
	TypeFloat32 = reflect.TypeOf(float32(0))
	TypeFloat64 = reflect.TypeOf(float64(0))
	TypeInt     = reflect.TypeOf(0)
	TypeInt8    = reflect.TypeOf(int8(0))
	TypeInt16   = reflect.TypeOf(int16(0))
	TypeInt32   = reflect.TypeOf(int32(0))
	TypeInt64   = reflect.TypeOf(int64(0))
	TypeUint    = reflect.TypeOf(uint(0))
	TypeUint8   = reflect.TypeOf(uint8(0))
	TypeUint16  = reflect.TypeOf(uint16(0))
	TypeUint32  = reflect.TypeOf(uint32(0))
	TypeUint64  = reflect.TypeOf(uint64(0))
)

func IsPrimitive(v interface{}) bool {
	switch v.(type) {
	case string, bool, float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	default:
		return false
	}
}

func FromString(s string, t reflect.Type) (interface{}, error) {
	if s == "" {
		return nil, nil
	}

	switch t {
	case TypeString:
		return s, nil
	case TypeBool:
		return strconv.ParseBool(s)
	case TypeFloat32:
		f, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return nil, err
		}
		return float32(f), nil
	case TypeFloat64:
		return strconv.ParseFloat(s, 64)
	case TypeInt:
		i, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, err
		}
		return int(i), nil
	case TypeInt8:
		i, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return nil, err
		}
		return int8(i), nil
	case TypeInt16:
		i, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return nil, err
		}
		return int16(i), nil
	case TypeInt32:
		i, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, err
		}
		return int32(i), nil
	case TypeInt64:
		return strconv.ParseInt(s, 10, 64)
	case TypeUint:
		i, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return nil, err
		}
		return uint(i), nil
	case TypeUint8:
		i, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return nil, err
		}
		return uint8(i), nil
	case TypeUint16:
		i, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return nil, err
		}
		return uint16(i), nil
	case TypeUint32:
		i, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return nil, err
		}
		return uint32(i), nil
	case TypeUint64:
		return strconv.ParseUint(s, 10, 64)
	default:
		return nil, fmt.Errorf("unsupported type: %v", t)
	}
}

// IsEmpty returns true if the slice is empty
func IsEmpty[T any](xs T) bool {
	switch v := reflect.TypeOf(xs); v.Kind() {
	case reflect.Map, reflect.Slice:
		return reflect.ValueOf(xs).Len() == 0
	default:
		return reflect.ValueOf(xs).IsZero()
	}
}

func IsNotEmpty[T any](xs T) bool {
	return !IsEmpty(xs)
}

// GetZero returns the zero value of the given type
func GetZero[T any]() T {
	var x T
	return x
}

func FromPointer[T any](p *T) T {
	if p == nil {
		return GetZero[T]()
	}

	return *p
}

func ToPointer[T any](v T) *T {
	if IsEmpty(v) {
		return nil
	}

	return &v
}
