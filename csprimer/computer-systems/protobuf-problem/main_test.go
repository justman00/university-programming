package main

import (
	"testing"
)

func FuzzMain(f *testing.F) {
	f.Add(uint64(73676))
	f.Fuzz(func(t *testing.T, orig uint64) {
		encoded := encode(orig)
		decoded := decode(encoded)
		t.Log(orig, encoded, decoded)
		if orig != decoded {
			t.Errorf("Before: %q, after: %q", orig, decoded)
		}
	})
}
