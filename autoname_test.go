package autoname

import (
	"strings"
	"testing"
)

func TestNameFormat(t *testing.T) {
	name := Generate()
	if !strings.Contains(name, "_") {
		t.Fatalf("Generated name does not contain an underscore")
	}
	if strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name contains numbers!")
	}
}

func BenchmarkGetRandomName(b *testing.B) {
	b.ReportAllocs()
	var out string
	for n := 0; n < b.N; n++ {
		out = Generate()
	}
	b.Log("Last result:", out)
}
