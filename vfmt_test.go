package vfmt

import (
	"fmt"
	"testing"
)

// vars is a list of variables used in the tests.
var vars = map[string]interface{}{
	"id":   12,
	"mode": "box",
}

// list is a list of matching input and expected output strings given the appropriate variables.
var list = []struct {
	src string
	out string
}{
	{"Hello, 世界", "Hello, 世界"},
	{"pic_{id|%04d}_{mode}.jpg", "pic_0012_box.jpg"},
	{"{mode}", "box"},
}

func TestSprintf(t *testing.T) {
	// test list
	for _, v := range list {
		out, err := Sprintf(v.src, vars)
		if err != nil {
			t.Fatal(err)
		}
		if out != v.out {
			t.Fatalf("expected %q, got %q\n", v.out, out)
		}
	}

	// test missing variables
	if _, err := Sprintf("foo{bar}", nil); err == nil {
		t.Fatal("expected missing variable error, got nil")
	}
	if _, err := Sprintf("foo{bar|%s}", nil); err == nil {
		t.Fatal("expected missing variable error, got nil")
	}

	// test template error
	if _, err := Sprintf("{x}", map[string]interface{}{"x": func() string { panic(0) }}); err == nil {
		t.Fatal("expected template error, got nil")
	}
}

func ExampleSprintf() {
	out, err := Sprintf("pic_{id|%04d}_{mode}.jpg", map[string]interface{}{
		"id":   12,
		"mode": "box",
	})
	if err != nil {
		// handle error, often a missing variable in the map
	}
	fmt.Println(out)
	// Output: pic_0012_box.jpg
}
