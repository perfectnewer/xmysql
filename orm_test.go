package xmysql

import (
	"testing"
)

type tUnmarshalModel struct {
	First  int    `mysql:"first"`
	Second string `mysql:"second"`
}

func Test_unmarshalResult(t *testing.T) {
	rst := make([]map[string]string, 0, 0)
	rst = append(rst, map[string]string{"first": "1", "second": "hello"})
	rst = append(rst, map[string]string{"first": "2", "second": "wawo"})
	out := make([]*tUnmarshalModel, 0, 0)
	if err := unmarshalResult(&out, rst); err != nil {
		t.Error(err)
	}
	if len(out) != 2 {
		t.Fail()
	}
	if out[0].First != 1 || out[0].Second != "hello" {
		t.Fail()
	}
	if out[1].First != 2 || out[1].Second != "wawo" {
		t.Fail()
	}

	tm := new(tUnmarshalModel)
	if err := unmarshalResult(tm, rst); err != nil {
		t.Error(err)
	}
}
