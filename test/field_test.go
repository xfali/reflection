// Copyright (C) 2019-2022, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/reflection"
	"testing"
	"time"
)

type testRootStruct struct {
	A int               `json:"a"`
	B testBranchStruct  `json:"b"`
	C *testBranchStruct `json:"c"`
}

type testBranchStruct struct {
	S string    `json:"s"`
	T time.Time `json:"t"`
	B bool      `json:"b"`
}

func TestField(t *testing.T) {
	c := testBranchStruct{}
	o := testRootStruct{
		A: -1,
		B: testBranchStruct{
			S: "BS1",
			B: false,
		},
		C: &c,
	}
	t.Run("without tag", func(t *testing.T) {
		err := reflection.SetStrcutFieldValue(o, "", 123)
		if err == nil {
			t.Fatal("cannot be here.", err)
		}

		err = reflection.SetStrcutFieldValue(&o, "A", 123)
		if err != nil {
			t.Fatal("cannot be here.", err)
		}
		if o.A != 123 {
			t.Fatal("expect 123 but get ", o.A)
		}

		err = reflection.SetStrcutFieldValue(&o, "B.S", "hello world")
		if err != nil {
			t.Fatal("cannot be here.", err)
		}
		if o.B.S != "hello world" {
			t.Fatal("expect hello world but get ", o.B.S)
		}

		err = reflection.SetStrcutFieldValue(&o, "B.B", "true")
		if err != nil {
			t.Fatal("cannot be here.", err)
		}
		if o.B.B != true {
			t.Fatal("expect true but get ", o.B.B)
		}

		err = reflection.SetStrcutFieldValue(&o, "C.T", "2022-07-01 19:00:00")
		if err != nil {
			t.Fatal("cannot be here.", err)
		}
		if o.C.T.Format("2006-01-02 15:04:05") != "2022-07-01 19:00:00" {
			t.Fatal("expect 2022-07-01 19:00:00 but get ", o.C.T.Format("2006-01-02 15:04:05"))
		}
	})

	t.Run("with tag", func(t *testing.T) {
		err := reflection.SetStrcutFieldValueByTag(o, "", 123, "json")
		if err == nil {
			t.Fatal("cannot be here.", err)
		}

		err = reflection.SetStrcutFieldValueByTag(&o, "a", 123, "json")
		if err != nil {
			t.Fatal("cannot be here.", err)
		}
		if o.A != 123 {
			t.Fatal("expect 123 but get ", o.A)
		}

		err = reflection.SetStrcutFieldValueByTag(&o, "b.s", "hello world", "json")
		if err != nil {
			t.Fatal("cannot be here.", err)
		}
		if o.B.S != "hello world" {
			t.Fatal("expect hello world but get ", o.B.S)
		}

		err = reflection.SetStrcutFieldValueByTag(&o, "b.b", true, "json")
		if err != nil {
			t.Fatal("cannot be here.", err)
		}
		if o.B.B != true {
			t.Fatal("expect true but get ", o.B.B)
		}

		err = reflection.SetStrcutFieldValueByTag(&o, "c.t", "2022-07-01 19:00:00", "json")
		if err != nil {
			t.Fatal("cannot be here.", err)
		}
		if o.C.T.Format("2006-01-02 15:04:05") != "2022-07-01 19:00:00" {
			t.Fatal("expect 2022-07-01 19:00:00 but get ", o.C.T.Format("2006-01-02 15:04:05"))
		}
	})
}
