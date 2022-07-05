// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package reflection

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func SafeSetValue(f reflect.Value, v reflect.Value) bool {
	if err := MustPtrValue(f); err != nil {
		return false
	}
	f = f.Elem()
	return SetValue(f, v)
}

func SetValue(f reflect.Value, vv reflect.Value) bool {
	hasAssigned := false
	rawValueType := vv.Type()

	ft := f.Type()
	switch ft.Kind() {
	case reflect.Bool:
		switch rawValueType.Kind() {
		case reflect.Bool:
			hasAssigned = true
			f.SetBool(vv.Bool())
			break
		case reflect.Slice:
			if d, ok := vv.Interface().([]uint8); ok {
				hasAssigned = true
				f.SetBool(d[0] != 0)
			}
			break
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
			hasAssigned = true
			f.SetBool(vv.Uint() != 0)
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			hasAssigned = true
			f.SetBool(vv.Int() != 0)
			break
		case reflect.String:
			b, err := strconv.ParseBool(vv.String())
			if err == nil {
				hasAssigned = true
				f.SetBool(b)
			}
			break
		}
		break
	case reflect.String:
		switch rawValueType.Kind() {
		case reflect.String:
			hasAssigned = true
			f.SetString(vv.String())
			break
		case reflect.Slice:
			if d, ok := vv.Interface().([]uint8); ok {
				hasAssigned = true
				f.SetString(string(d))
			}
			break
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
			hasAssigned = true
			f.SetString(strconv.FormatUint(vv.Uint(), 10))
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			hasAssigned = true
			f.SetString(strconv.FormatInt(vv.Int(), 10))
			break
		case reflect.Float64:
			hasAssigned = true
			f.SetString(strconv.FormatFloat(vv.Float(), 'g', -1, 64))
			break
		case reflect.Float32:
			hasAssigned = true
			f.SetString(strconv.FormatFloat(vv.Float(), 'g', -1, 32))
			break
		case reflect.Bool:
			hasAssigned = true
			f.SetString(strconv.FormatBool(vv.Bool()))
			break
		//case reflect.Struct:
		//    if ti, ok := v.(time.Time); ok {
		//        hasAssigned = true
		//        if ti.IsZero() {
		//            f.SetString("")
		//        } else {
		//            f.SetString(ti.String())
		//        }
		//    } else {
		//        hasAssigned = true
		//        f.SetString(fmt.Sprintf("%v", v))
		//    }
		default:
			hasAssigned = true
			f.SetString(fmt.Sprintf("%v", vv.Interface()))
		}
		break
	case reflect.Complex64, reflect.Complex128:
		switch rawValueType.Kind() {
		case reflect.Complex64, reflect.Complex128:
			hasAssigned = true
			f.SetComplex(vv.Complex())
			break
		case reflect.Slice:
			if rawValueType.ConvertibleTo(BytesType) {
				d := vv.Bytes()
				if len(d) > 0 {
					if f.CanAddr() {
						err := json.Unmarshal(d, f.Addr().Interface())
						if err != nil {
							return false
						}
					} else {
						x := reflect.New(ft)
						err := json.Unmarshal(d, x.Interface())
						if err != nil {
							return false
						}
						hasAssigned = true
						f.Set(x.Elem())
						break
					}
				}
			}
			break
		}
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch rawValueType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			hasAssigned = true
			f.SetInt(vv.Int())
			break
		case reflect.Slice:
			if d, ok := vv.Interface().([]uint8); ok {
				intV, err := strconv.ParseInt(string(d), 10, 64)
				if err == nil {
					hasAssigned = true
					f.SetInt(intV)
				}
			}
			break
		case reflect.String:
			b, err := strconv.ParseInt(vv.String(), 10, 64)
			if err == nil {
				hasAssigned = true
				f.SetInt(b)
			}
			break
		}
		break
	case reflect.Float32, reflect.Float64:
		switch rawValueType.Kind() {
		case reflect.Float32, reflect.Float64:
			hasAssigned = true
			f.SetFloat(vv.Float())
			break
		case reflect.Slice:
			if d, ok := vv.Interface().([]uint8); ok {
				floatV, err := strconv.ParseFloat(string(d), 64)
				if err == nil {
					hasAssigned = true
					f.SetFloat(floatV)
				}
			}
			break
		case reflect.String:
			b, err := strconv.ParseFloat(vv.String(), 10)
			if err == nil {
				hasAssigned = true
				f.SetFloat(b)
			}
			break
		}
		break
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		switch rawValueType.Kind() {
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
			hasAssigned = true
			f.SetUint(vv.Uint())
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			hasAssigned = true
			f.SetUint(uint64(vv.Int()))
			break
		case reflect.Slice:
			if d, ok := vv.Interface().([]uint8); ok {
				uintV, err := strconv.ParseUint(string(d), 10, 64)
				if err == nil {
					hasAssigned = true
					f.SetUint(uintV)
				}
			}
			break
		case reflect.String:
			b, err := strconv.ParseUint(vv.String(), 10, 64)
			if err == nil {
				hasAssigned = true
				f.SetUint(b)
			}
			break
		}
		break
	case reflect.Struct:
		fieldType := f.Type()
		if fieldType.ConvertibleTo(TimeType) {
			if rawValueType == TimeType {
				hasAssigned = true
				t := vv.Convert(TimeType).Interface().(time.Time)
				f.Set(reflect.ValueOf(t).Convert(fieldType))
			} else if rawValueType == IntType || rawValueType == Int64Type ||
				rawValueType == Int32Type {
				hasAssigned = true

				t := time.Unix(vv.Int(), 0)
				f.Set(reflect.ValueOf(t).Convert(fieldType))
			} else if rawValueType == StringType {
				t, err := convert2Time([]byte(vv.String()), time.Local)
				if err == nil {
					hasAssigned = true
					f.Set(reflect.ValueOf(t).Convert(fieldType))
				}
			} else {
				if d, ok := vv.Interface().([]byte); ok {
					t, err := convert2Time(d, time.Local)
					if err == nil {
						hasAssigned = true
						f.Set(reflect.ValueOf(t).Convert(fieldType))
					}
				}
			}
		} else {
			f.Set(vv)
		}
		break
	case reflect.Interface:
		hasAssigned = true
		f.Set(vv)
		break
	}

	return hasAssigned
}

const (
	zeroTime0 = "0000-00-00 00:00:00"
	zeroTime1 = "0001-01-01 00:00:00"
)

func convert2Time(data []byte, location *time.Location) (time.Time, error) {
	timeStr := strings.TrimSpace(string(data))
	var timeRet time.Time
	var err error
	if timeStr == zeroTime0 || timeStr == zeroTime1 {
	} else if !strings.ContainsAny(timeStr, "- :") {
		// time stamp
		sd, err := strconv.ParseInt(timeStr, 10, 64)
		if err == nil {
			timeRet = time.Unix(sd, 0)
		}
	} else if len(timeStr) > 19 && strings.Contains(timeStr, "-") {
		timeRet, err = time.ParseInLocation(time.RFC3339Nano, timeStr, location)
		if err != nil {
			timeRet, err = time.ParseInLocation("2006-01-02 15:04:05.999999999", timeStr, location)
		}
		if err != nil {
			timeRet, err = time.ParseInLocation("2006-01-02 15:04:05.9999999 Z07:00", timeStr, location)
		}
	} else if len(timeStr) == 19 && strings.Contains(timeStr, "-") {
		timeRet, err = time.ParseInLocation("2006-01-02 15:04:05", timeStr, location)
	} else if len(timeStr) == 10 && timeStr[4] == '-' && timeStr[7] == '-' {
		timeRet, err = time.ParseInLocation("2006-01-02", timeStr, location)
	}
	return timeRet, nil
}

func MustPtr(bean interface{}) error {
	return MustPtrValue(reflect.ValueOf(bean))
}

func MustPtrValue(beanValue reflect.Value) error {
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("Object is not a pointer. ")
	}
	//else if beanValue.Elem().Kind() == reflect.Ptr {
	//	return errors.New("Object's value is a pointer. ")
	//}
	return nil
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	return CheckValueNilSafe(reflect.ValueOf(i))
}

func CheckValueNilSafe(vi reflect.Value) bool {
	switch vi.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return vi.IsNil()
	}
	return false
}

func CanSet(i interface{}) bool {
	if i == nil {
		return false
	}
	vi := reflect.ValueOf(i)
	if MustPtrValue(vi) != nil {
		return false
	}
	if vi.IsNil() {
		return false
	}
	return true
}

func NewValue(t reflect.Type) reflect.Value {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Ptr {
		panic("Error Type")
	}
	return reflect.New(t).Elem()
}

func New(t reflect.Type) interface{} {
	return NewValue(t).Interface()
}
