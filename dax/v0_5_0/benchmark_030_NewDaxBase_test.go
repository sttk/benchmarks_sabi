package v0_5_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
)

func BenchmarkDax_____NewDaxBase(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base := sabi.NewDaxBase()
		_ = base
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_NewDaxBase(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		base := prev.NewDaxBase()
		_ = base
	}
	b.StopTimer()
}
