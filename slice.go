// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package reflection

import (
	"errors"
	"reflect"
)

func CopySliceInterface(dest, src interface{}) (int, error) {
	t := reflect.TypeOf(dest)
	v := reflect.ValueOf(dest)
	if t.Kind() != reflect.Ptr {
		return 0, errors.New("Dest Type is not a ptr. " + t.String())
	}

	return CopySlice(v.Elem(), reflect.ValueOf(src))
}

func CopySlice(dest, src reflect.Value) (int, error) {
	destType := dest.Type()
	if destType.Kind() != reflect.Slice {
		return 0, errors.New("Dest Type is not a slice ptr. " + destType.String())
	}
	destElemType := destType.Elem()
	srcType := src.Type()
	if srcType.Kind() != reflect.Slice {
		return 0, errors.New("Src Type is not a slice. " + srcType.String())
	}
	srcElemType := srcType.Elem()

	if destElemType.Kind() == srcElemType.Kind() {
		dest.Set(src)
		return dest.Len(), nil
	} else {
		n := 0
		destTmp := dest
		for i := 0; i < src.Len(); i++ {
			ov := src.Index(i)
			ot := ov.Type()
			// interface
			if ot.ConvertibleTo(destElemType) {
				ov = ov.Convert(destElemType)
				ot = ov.Type()
			}
			if ot.AssignableTo(destElemType) {
				destTmp = reflect.Append(destTmp, ov)
				n++
			}
		}
		dest.Set(destTmp)
		return n, nil
	}
}
