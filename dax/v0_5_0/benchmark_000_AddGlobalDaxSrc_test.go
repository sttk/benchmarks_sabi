package v0_5_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_5_0"

	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_4_0/supp"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
)

func BenchmarkDax_____AddGlobalDaxSrc_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.ResetGlobals()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_AddGlobalDaxSrc_zeroDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.ResetGlobals()
	}
	b.StopTimer()
}

func BenchmarkDax_____AddGlobalDaxSrc_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})
		sabi.ResetGlobals()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_AddGlobalDaxSrc_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
		prev.ResetGlobals()
	}
	b.StopTimer()
}

func BenchmarkDax_____AddGlobalDaxSrc_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})
		sabi.AddGlobalDaxSrc("database", supp.FooDaxSrc{})
		sabi.AddGlobalDaxSrc("pubsub", supp.FooDaxSrc{})
		sabi.AddGlobalDaxSrc("json", supp.FooDaxSrc{})
		sabi.AddGlobalDaxSrc("env", supp.FooDaxSrc{})

		sabi.ResetGlobals()
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_AddGlobalDaxSrc_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
		prev.AddGlobalDaxSrc("database", prev_supp.FooDaxSrc{})
		prev.AddGlobalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
		prev.AddGlobalDaxSrc("json", prev_supp.FooDaxSrc{})
		prev.AddGlobalDaxSrc("env", prev_supp.FooDaxSrc{})

		prev.ResetGlobals()
	}
	b.StopTimer()
}
