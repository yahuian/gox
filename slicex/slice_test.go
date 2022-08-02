package slicex_test

import (
	"reflect"
	"sort"
	"strconv"
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

func TestMapFilterReduce(t *testing.T) {
	list := []string{"1", "2", "3", "4", "5"}

	// map
	mapRes := slicex.Map(list, func(index int, value string) int {
		i, err := strconv.Atoi(value)
		if err != nil {
			t.Fatal(err)
		}
		return i
	})
	mapWant := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(mapWant, mapRes) {
		t.Errorf("want %+v get %+v", mapWant, mapRes)
	}

	// filter
	filterRes := slicex.Filter(mapRes, func(index int, value int) bool {
		return value%2 == 0
	})
	filterWant := []int{2, 4}
	if !reflect.DeepEqual(filterWant, filterRes) {
		t.Errorf("want %+v get %+v", filterWant, filterRes)
	}

	// reduce
	reduceRes := slicex.Reduce(filterRes, 1, func(index, result, value int) int {
		return result + value
	})
	if !reflect.DeepEqual(7, reduceRes) {
		t.Errorf("want %+v get %+v", 7, reduceRes)
	}
}
