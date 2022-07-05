// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/reflection"
	"testing"
)


func TestCopyMap(t *testing.T) {
	t.Run("not map", func(t *testing.T) {
		src := 0
		dst := map[int]int{}
		n, err := reflection.CopyMapInterface(&dst, src)
		if err == nil {
			t.Fatal("Must cannot copy!")
		}
		t.Log(n, " ", err)

		n, err = reflection.CopyMapInterface(src, dst)
		if err == nil {
			t.Fatal("Must cannot copy!")
		}
		t.Log(n, " ", err)
	})

	t.Run("map int int", func(t *testing.T) {
		src := map[int]int{0:0, 1:1, 2:2}
		var dst map[int]int
		n, err := reflection.CopyMapInterface(&dst, src)
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

	t.Run("map int int64", func(t *testing.T) {
		src := map[int]int{0:0, 1:1, 2:2}
		var dst map[int]int64
		n, err := reflection.CopyMapInterface(&dst, src)
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

	t.Run("map testStr String", func(t *testing.T) {
		src := map[string]testStr{"hello":"hello", "world":"world"}
		var dst  map[string]testStr
		n, err := reflection.CopyMapInterface(&dst, src)
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
