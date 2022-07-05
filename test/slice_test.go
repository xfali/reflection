// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"fmt"
	"github.com/xfali/reflection"
	"reflect"
	"testing"
)

type testStr string

func (t testStr) String() string {
	return string(t)
}

func TestCopySlice(t *testing.T) {
	t.Run("not slice", func(t *testing.T) {
		src := 0
		dst := []int{0}
		n, err := reflection.CopySlice(reflect.ValueOf(src), reflect.ValueOf(dst))
		if err == nil {
			t.Fatal("Must cannot copy!")
		}
		t.Log(n, " ", err)

		n, err = reflection.CopySlice(reflect.ValueOf(dst), reflect.ValueOf(src))
		if err == nil {
			t.Fatal("Must cannot copy!")
		}
		t.Log(n, " ", err)
	})

	t.Run("slice int int", func(t *testing.T) {
		src := []int{0, 1, 2}
		var dst []int
		n, err := reflection.CopySlice(reflect.ValueOf(&dst), reflect.ValueOf(src))
		if err != nil {
			t.Fatal("Must copy!", err)
		}
		if n != 3 {
			t.Fatal("expect 3 get  ", n)
		}

		for i := range src {
			if src[i] != dst[i] {
				t.Fatalf("expect %d get %d ", src[i], dst[i])
			}
		}
	})

	t.Run("slice int int64", func(t *testing.T) {
		src := []int{0, 1, 2}
		var dst []int64
		n, err := reflection.CopySlice(reflect.ValueOf(&dst), reflect.ValueOf(src))
		if err != nil {
			t.Fatal("Must copy!", err)
		}
		if n != 3 {
			t.Fatal("expect 3 get  ", n)
		}

		for i := range src {
			if int64(src[i]) != dst[i] {
				t.Fatalf("expect %d get %d ", src[i], dst[i])
			}
		}
	})

	t.Run("slice testStr String", func(t *testing.T) {
		src := []testStr{"hello", "world"}
		var dst []fmt.Stringer
		n, err := reflection.CopySlice(reflect.ValueOf(&dst), reflect.ValueOf(src))
		if err != nil {
			t.Fatal("Must copy!", err)
		}
		if n != 2 {
			t.Fatal("expect 2 get  ", n)
		}

		for i := range src {
			if src[i].String() != dst[i].String() {
				t.Fatalf("expect %s get %s ", src[i], dst[i])
			} else {
				t.Log(dst[i])
			}
		}
	})
}
