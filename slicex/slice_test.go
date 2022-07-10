package slicex_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/yahuian/gox/slicex"
)

func TestUnique(t *testing.T) {
	list := []int{2, 1, 2, 2, 3, 5}
	res := slicex.Unique(list)
	sort.Ints(res)
	want := []int{1, 2, 3, 5}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("want %+v get %+v", want, res)
	}
}

func TestContain(t *testing.T) {
	list := []string{"python", "java", "go"}
	if !slicex.Contain(list, "go") {
		t.Errorf("want %+v get %v", true, false)
	}
	if slicex.Contain(list, "golang") {
		t.Errorf("want %+v get %v", false, true)
	}
}
