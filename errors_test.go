package errors

import (
	"errors"
	"testing"
)

func TestErr(t *testing.T) {
	e := New(UnknowErr, "a err")
	t.Log(e)

	e = Wrap(UnknowErr, e, "err")
	t.Log(e)
}

func TestType(t *testing.T) {
	e := errors.New("err")
	t.Log(Type(e))

	e = New(2, "err type 2")
	e = WithMsg(e, "with msg")
	t.Log(e, Type(e))
	if Type(e) != 2 {
		t.Fail()
	}

	e = New(-1, "unknow err")
	e = Wrap(1, e, "err type 1")
	t.Log(e, Type(e))
	if Type(e) != 1 { //取最外层的type
		t.Fail()
	}
}
