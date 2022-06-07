package autoname

import (
	"strings"
	"testing"
)

const delimiter = "-"

func TestNameFormat(t *testing.T) {
	name := Generate(delimiter)
	if !strings.Contains(name, delimiter) {
		t.Fatalf("Generated name does not contain correct delimiter")
	}
	if strings.ContainsAny(name, "0123456789") {
		t.Fatalf("Generated name contains numbers!")
	}
}

func BenchmarkGetRandomName(b *testing.B) {
	b.ReportAllocs()
	var out string
	for n := 0; n < b.N; n++ {
		out = Generate(delimiter)
	}
	b.Log("Last result:", out)
}
