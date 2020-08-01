package mruby

import (
	"testing"
)

func TestHash(t *testing.T) {
	mrb := NewMrb()
	defer mrb.Close()

	value, err := mrb.LoadString(`{"foo" => "bar", "baz" => false, :fizz => 1, buzz: 2.0}`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	h := value.Hash()

	// Get
	value, err = h.Get(String("foo"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if value.String() != "bar" {
		t.Fatalf("bad: %s", value)
	}

	// Get false type
	value, err = h.Get(String("baz"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if valType := value.Type(); valType != TypeFalse {
		t.Fatalf("bad type: %v", valType)
	}
	if value.String() != "false" {
		t.Fatalf("bad: %s", value)
	}

	// Get fixnum type with symbol key
	value, err = h.Get(Symbol("fizz"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if valType := value.Type(); valType != TypeFixnum {
		t.Fatalf("bad type: %v", valType)
	}
	if value.Fixnum() != 1 {
		t.Fatalf("bad: %s", value)
	}

	// Get float type with symbol key
	value, err = h.Get(Symbol("buzz"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if valType := value.Type(); valType != TypeFloat {
		t.Fatalf("bad type: %v", valType)
	}
	if value.Float() != 2.0 {
		t.Fatalf("bad: %s", value)
	}

	// Set
	err = h.Set(String("foo"), String("baz"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	value, err = h.Get(String("foo"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if value.String() != "baz" {
		t.Fatalf("bad: %s", value)
	}

	// Keys
	value, err = h.Keys()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if value.Type() != TypeArray {
		t.Fatalf("bad: %v", value.Type())
	}
	if value.String() != `["foo", "baz", :fizz, :buzz]` {
		t.Fatalf("bad: %s", value)
	}

	// Delete
	value, err = h.Delete(String("foo"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if value.String() != "baz" {
		t.Fatalf("bad: %s", value)
	}

	value, err = h.Keys()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if value.String() != `["baz", :fizz, :buzz]` {
		t.Fatalf("bad: %s", value)
	}

	// Delete non-existing
	value, err = h.Delete(String("nope"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if value != nil {
		t.Fatalf("bad: %s", value)
	}
}
