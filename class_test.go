package mruby

import (
	"testing"
)

func TestClassDefineClassMethod(t *testing.T) {
	mrb := NewMrb()
	defer mrb.Close()

	class := mrb.DefineClass("Hello", mrb.ObjectClass())
	class.DefineClassMethod("foo", testCallback, ArgsNone())
	value, err := mrb.LoadString("Hello.foo")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	testCallbackResult(t, value)
}

func TestClassDefineMethod(t *testing.T) {
	mrb := NewMrb()
	defer mrb.Close()

	class := mrb.DefineClass("Hello", mrb.ObjectClass())
	class.DefineMethod("foo", testCallback, ArgsNone())
	value, err := mrb.LoadString("Hello.new.foo")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	testCallbackResult(t, value)
}

func TestClassValue(t *testing.T) {
	mrb := NewMrb()
	defer mrb.Close()

	class := mrb.DefineClass("Hello", mrb.ObjectClass())
	value := class.MrbValue()
	if value.Type() != TypeClass {
		t.Fatalf("bad: %d", value.Type())
	}
}
