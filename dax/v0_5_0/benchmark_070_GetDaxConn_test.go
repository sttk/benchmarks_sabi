package v0_5_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_5_0"

	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_4_0/supp"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
)

func BenchmarkDax_____GetDaxConn_global_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})

	base := sabi.NewDaxBase()

	sabi.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_GetDaxConn_global_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})

	base := prev.NewDaxBase()

	prev.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := prev.GetDaxConn[prev_supp.FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func BenchmarkDax_____GetDaxConn_global_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("database", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("pubsub", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("json", supp.FooDaxSrc{})
	sabi.AddGlobalDaxSrc("env", supp.FooDaxSrc{})

	base := sabi.NewDaxBase()

	sabi.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err0 := sabi.GetDaxConn[supp.FooDaxConn](base, "cliargs")
		conn1, err1 := sabi.GetDaxConn[supp.FooDaxConn](base, "database")
		conn2, err2 := sabi.GetDaxConn[supp.FooDaxConn](base, "pubsub")
		conn3, err3 := sabi.GetDaxConn[supp.FooDaxConn](base, "json")
		conn4, err4 := sabi.GetDaxConn[supp.FooDaxConn](base, "env")
		_ = conn0
		_ = conn1
		_ = conn2
		_ = conn3
		_ = conn4
		_ = err0
		_ = err1
		_ = err2
		_ = err3
		_ = err4
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_GetDaxConn_global_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	prev.AddGlobalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("database", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("json", prev_supp.FooDaxSrc{})
	prev.AddGlobalDaxSrc("env", prev_supp.FooDaxSrc{})

	base := prev.NewDaxBase()

	prev.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err0 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "cliargs")
		conn1, err1 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "database")
		conn2, err2 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "pubsub")
		conn3, err3 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "json")
		conn4, err4 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "env")
		_ = conn0
		_ = conn1
		_ = conn2
		_ = conn3
		_ = conn4
		_ = err0
		_ = err1
		_ = err2
		_ = err3
		_ = err4
	}
	b.StopTimer()
}

func BenchmarkDax_____GetDaxConn_local_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.AddLocalDaxSrc("cliargs", supp.FooDaxSrc{})

	sabi.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_GetDaxConn_local_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})

	prev.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn, err := prev.GetDaxConn[prev_supp.FooDaxConn](base, "cliargs")
		_ = conn
		_ = err
	}
	b.StopTimer()
}

func BenchmarkDax_____GetDaxConn_local_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()

	base.AddLocalDaxSrc("cliargs", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("database", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("pubsub", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("json", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("env", supp.FooDaxSrc{})

	sabi.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err0 := sabi.GetDaxConn[supp.FooDaxConn](base, "cliargs")
		conn1, err1 := sabi.GetDaxConn[supp.FooDaxConn](base, "database")
		conn2, err2 := sabi.GetDaxConn[supp.FooDaxConn](base, "pubsub")
		conn3, err3 := sabi.GetDaxConn[supp.FooDaxConn](base, "json")
		conn4, err4 := sabi.GetDaxConn[supp.FooDaxConn](base, "env")
		_ = conn0
		_ = conn1
		_ = conn2
		_ = conn3
		_ = conn4
		_ = err0
		_ = err1
		_ = err2
		_ = err3
		_ = err4
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_GetDaxConn_local_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()

	base.SetUpLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("database", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("json", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("env", prev_supp.FooDaxSrc{})

	prev.Begin(base)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		conn0, err0 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "cliargs")
		conn1, err1 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "database")
		conn2, err2 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "pubsub")
		conn3, err3 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "json")
		conn4, err4 := prev.GetDaxConn[prev_supp.FooDaxConn](base, "env")
		_ = conn0
		_ = conn1
		_ = conn2
		_ = conn3
		_ = conn4
		_ = err0
		_ = err1
		_ = err2
		_ = err3
		_ = err4
	}
	b.StopTimer()
}
