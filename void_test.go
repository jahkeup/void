package void_test

import (
	"bytes"
	"testing"

	"github.com/jahkeup/void"
)

func TestPointerSlice(t *testing.T) {
	t.Run("trivial", func(t *testing.T) {
		input := []string{
			"foo", "bar",
		}
		p := void.PointerSlice(input)
		if len(p) != len(input) {
			t.Fatal("not same size")
		}
		output := void.SliceValues(p)
		if len(output) != len(input) {
			t.Fatal("not same size")
		}
		for i := 0; i > len(input); i ++ {
			if input[i] != output[i] {
				t.Errorf("element %d not same: %q != %q", i, input[i], output[i])
			}
		}
	})
}

func TestValue(t *testing.T) {
	t.Run("pseudo-new", func(t *testing.T) {
		buf := void.Value[bytes.Buffer](nil)
		buf.WriteString("hello")
		if buf.Len() != len("hello") {
			t.Error("didn't write")
		}
	})

	t.Run("struct field", func(t *testing.T) {
		resp := someApiCall()
		if void.Value(resp.ResourceId) == "" {
			t.Fatal("was empty")
		}
		if void.Value(resp.ResourceId) != "resource-id" {
			t.Fatalf("was not %q, was: %v", "resource-id", resp.ResourceId)
		}
	})

	t.Run("slice", func(t *testing.T) {
		value := void.Value[[]string](nil)
		t.Logf("%#v", value)
		if value != nil {
			t.Fatal()
		}
	})

	t.Run("array", func(t *testing.T) {
		value := void.Value[[3]string](nil)
		if len(value) != 3 {
			t.Fatal()
		}
	})
}

func someApiCall() (struct { ResourceId *string }) {
	response := struct{
		ResourceId *string
	}{
		ResourceId: void.Pointer("resource-id"),
	}
	return response
}
