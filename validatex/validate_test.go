package validatex_test

import (
	"testing"

	"github.com/gin-gonic/gin/binding"
	"github.com/yahuian/gox/validatex"
)

type user struct {
	Name string `validate:"required"`
	Age  int    `validate:"gt=18"`
}

func TestStruct(t *testing.T) {
	if err := validatex.Init(); err != nil {
		t.Fatal(err)
	}

	var u user
	var want = "Name is a required field, Age must be greater than 18"
	msg := validatex.Struct(u).Error()
	if msg != want {
		t.Errorf("want `%s` but get `%s`", want, msg)
	}
}

func TestVar(t *testing.T) {
	if err := validatex.Init(); err != nil {
		t.Fatal(err)
	}

	u := "tom"
	var want = "user must be at least 5 characters in length"
	msg := validatex.Val(u, "user", "min=5").Error()
	if msg != want {
		t.Errorf("want `%s` but get `%s`", want, msg)
	}
}

func TestGin(t *testing.T) {
	if err := validatex.Init(validatex.WithGin()); err != nil {
		t.Fatal(err)
	}

	var u user
	u.Age = 19
	var want = "Name is a required field"
	msg := binding.JSON.BindBody([]byte(`{}`), &u).Error()
	if msg != want {
		t.Errorf("want `%s` but get `%s`", want, msg)
	}
}
