/*
 * Copyright 2023 Xiongfa Li.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"github.com/xfali/reflection"
	"reflect"
	"testing"
	"time"
)

type TestTable struct {
	Id       int64  `column:"id"`
	Username string `column:"username"`
	Password string `column:"password"`
}

func TestReflectObjectStruct(t *testing.T) {
	v := TestTable{}
	info, err := reflection.GetObjectInfo(&v)
	if err != nil {
		t.Fatal()
	}
	t.Logf("classname :%v", info.GetClassName())
	t.Log(v)
	newOne := TestTable{
		Username: "1",
	}
	info.SetValue(reflect.ValueOf(newOne))

	if v.Username != "1" {
		t.Fatal("Expect 1 but get ", v.Username)
	}
	t.Logf("after set :%v\n", v)

	info.SetField("username", reflect.ValueOf("123"))
	if v.Username != "123" {
		t.Fatal("Expect 123 but get ", v.Username)
	}
	t.Logf("after setField :%v\n", v)
}

func TestReflectObjectSimpleTime(t *testing.T) {
	v := time.Time{}
	info, err := reflection.GetObjectInfo(&v)
	if err != nil {
		t.Fatal()
	}
	t.Logf("classname :%v", info.GetClassName())
	t.Log(v)
	newOne := TestTable{
		Username: "1",
	}
	info.SetValue(reflect.ValueOf(newOne))

	t.Logf("after set error type :%v\n", v)

	now := time.Now()
	info.SetValue(reflect.ValueOf(now))

	if !v.Equal(now) {
		t.Fatalf("expect %v but get ", v)
	}
	t.Logf("after set now type :%v\n", v)

	info.SetField("username", reflect.ValueOf("123"))
	t.Logf("after setField :%v\n", v)
}

func TestReflectObjectSimpleFloat(t *testing.T) {
	v := 0.0
	info, err := reflection.GetObjectInfo(&v)
	if err != nil {
		t.Fatal()
	}
	t.Logf("classname :%v", info.GetClassName())
	t.Log(v)
	info.SetValue(reflect.ValueOf(1))

	// 1 is int not float
	if v != 0 {
		t.Fatal("Expect 1 but get ", v)
	}
	t.Logf("after set int type :%v\n", v)

	info.SetValue(reflect.ValueOf(1.5))

	if v != 1.5 {
		t.Fatal("Expect 1.5 but get ", v)
	}

	t.Logf("after set float type :%v\n", v)
}

func TestReflectObjectMap(t *testing.T) {
	v := map[string]interface{}{}
	info, err := reflection.GetObjectInfo(&v)
	if err != nil {
		t.Fatal()
	}
	t.Logf("classname :%v", info.GetClassName())
	t.Log(v)
	info.SetValue(reflect.ValueOf(1))

	t.Logf("after set int type :%v\n", v)

	info.SetValue(reflect.ValueOf(map[string]int{"1": 1, "2": 2}))

	if v["1"] == 1 {
		t.Fatal("Expect nil but get ", v["1"])
	}
	t.Logf("after set map[string]int type :%v\n", v)

	info.SetValue(reflect.ValueOf(map[string]interface{}{"1": 1, "2": 3}))

	if v["1"] != 1 {
		t.Fatal("Expect 1 but get ", v["1"])
	}
	if v["2"] != 3 {
		t.Fatal("Expect 3 but get ", v["3"])
	}
	t.Logf("after set map[string]interface{} type :%v\n", v)

	info.SetField("username", reflect.ValueOf("123"))
	if v["username"] != "123" {
		t.Fatal("Expect 123 but get ", v["username"])
	}
	t.Logf("after setField username 123 :%v\n", v)

	info.SetField("username", reflect.ValueOf("321"))
	if v["username"] != "321" {
		t.Fatal("Expect 321 but get ", v["username"])
	}
	t.Logf("after setField username 321 :%v\n", v)
}

func TestReflectObjectSlice2(t *testing.T) {
	v := []int{}
	info, err := reflection.GetObjectInfo(&v)
	if err != nil {
		t.Fatal()
	}
	t.Logf("classname :%v", info.GetClassName())
	t.Log(v)
	info.SetValue(reflect.ValueOf(1))

	t.Logf("after set int type :%v\n", v)

	info.SetValue(reflect.ValueOf([]float32{1.0, 2, 3}))

	t.Logf("after set []float32{1.0,2,3} :%v\n", v)

	info.SetValue(reflect.ValueOf([]int{1, 2, 3}))

	if v[1] != 2 {
		t.Fatal("Expect 2 but get ", v[1])
	}
	t.Logf("after set []int{1,2,3} :%v\n", v)

	info.SetField("username", reflect.ValueOf(123))
	t.Logf("after setField :%v\n", v)

	info.AddValue(reflect.ValueOf(123))
	if v[3] != 123 {
		t.Fatal("Expect 123 but get ", v[3])
	}
	t.Logf("after AddValue :%v\n", v)
}

func TestReflectObjectSlice(t *testing.T) {
	v := []TestTable{}
	info, err := reflection.GetObjectInfo(&v)
	if err != nil {
		t.Fatal()
	}
	t.Logf("classname :%v", info.GetClassName())
	t.Log(v)
	info.SetValue(reflect.ValueOf(1))

	t.Logf("after set int type :%v\n", v)

	info.SetValue(reflect.ValueOf([]float32{1.0, 2, 3}))

	t.Logf("after set []float32{1.0,2,3} :%v\n", v)

	info.SetValue(reflect.ValueOf([]TestTable{{Username: "1", Password: "1"}}))

	if v[0].Username != "1" {
		t.Fatal("Expect 1 but get ", v[0].Username)
	}
	t.Logf(`after set []TestTable{{Username:"1", Password:"1"}} :%v\n`, v)

	info.SetField("username", reflect.ValueOf(123))
	t.Logf("after setField :%v\n", v)

	info.AddValue(reflect.ValueOf(1))
	t.Logf("after AddValue 1 :%v\n", v)

	info.AddValue(reflect.ValueOf(TestTable{Username: "2", Password: "2"}))
	if v[1].Username != "2" {
		t.Fatal("Expect 2 but get ", v[0].Username)
	}
	t.Logf(`after AddValue TestTable{Username: "2", Password:"2"} :%v\n`, v)

	ev := info.NewElem()
	ev.SetField("username", reflect.ValueOf("x"))
	info.AddValue(ev.GetValue())
	if v[2].Username != "x" {
		t.Fatal("Expect x but get ", v[1].Username)
	}

	t.Logf(`after AddValue new elem {Username: "x"} :%v\n`, v)
}
