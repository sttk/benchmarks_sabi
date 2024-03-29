package v0_5_0_test

import (
	"testing"

	prev "github.com/sttk/benchmarks_sabi/dax/v0_4_0"
	sabi "github.com/sttk/benchmarks_sabi/dax/v0_5_0"

	prev_supp "github.com/sttk/benchmarks_sabi/dax/v0_4_0/supp"
	supp "github.com/sttk/benchmarks_sabi/dax/v0_5_0/supp"
)

func BenchmarkDax_____RunTxn_commit_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.AddLocalDaxSrc("cliargs", supp.FooDaxSrc{})

	logic := func(dax sabi.Dax) sabi.Err {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](dax, "cliargs")
		_ = conn
		_ = err
		return sabi.Ok()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_RunTxn_commit_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})

	logic := func(dax prev.Dax) prev.Err {
		conn, err := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "cliargs")
		_ = conn
		_ = err
		return prev.Ok()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.RunTxn[prev.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDax_____RunTxn_commit_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.AddLocalDaxSrc("cliargs", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("database", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("pubsub", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("json", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("env", supp.FooDaxSrc{})

	logic := func(dax sabi.Dax) sabi.Err {
		conn0, err0 := sabi.GetDaxConn[supp.FooDaxConn](dax, "cliargs")
		conn1, err1 := sabi.GetDaxConn[supp.FooDaxConn](dax, "database")
		conn2, err2 := sabi.GetDaxConn[supp.FooDaxConn](dax, "pubsub")
		conn3, err3 := sabi.GetDaxConn[supp.FooDaxConn](dax, "json")
		conn4, err4 := sabi.GetDaxConn[supp.FooDaxConn](dax, "env")
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
		return sabi.Ok()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_RunTxn_commit_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("database", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("json", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("env", prev_supp.FooDaxSrc{})

	logic := func(dax prev.Dax) prev.Err {
		conn0, err0 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "cliargs")
		conn1, err1 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "database")
		conn2, err2 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "pubsub")
		conn3, err3 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "json")
		conn4, err4 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "env")
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
		return prev.Ok()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.RunTxn[prev.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDax_____RunTxn_rollback_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.AddLocalDaxSrc("cliargs", supp.FooDaxSrc{})

	type Fail struct{}

	logic := func(dax sabi.Dax) sabi.Err {
		conn, err := sabi.GetDaxConn[supp.FooDaxConn](dax, "cliargs")
		_ = conn
		_ = err
		return sabi.NewErr(Fail{})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_RunTxn_rollback_oneDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})

	type Fail struct{}

	logic := func(dax prev.Dax) prev.Err {
		conn, err := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "cliargs")
		_ = conn
		_ = err
		return prev.NewErr(Fail{})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.RunTxn[prev.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDax_____RunTxn_rollback_fiveDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	base := sabi.NewDaxBase()
	base.AddLocalDaxSrc("cliargs", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("database", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("pubsub", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("json", supp.FooDaxSrc{})
	base.AddLocalDaxSrc("env", supp.FooDaxSrc{})

	type Fail struct{}

	logic := func(dax sabi.Dax) sabi.Err {
		conn0, err0 := sabi.GetDaxConn[supp.FooDaxConn](dax, "cliargs")
		conn1, err1 := sabi.GetDaxConn[supp.FooDaxConn](dax, "database")
		conn2, err2 := sabi.GetDaxConn[supp.FooDaxConn](dax, "pubsub")
		conn3, err3 := sabi.GetDaxConn[supp.FooDaxConn](dax, "json")
		conn4, err4 := sabi.GetDaxConn[supp.FooDaxConn](dax, "env")
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
		return sabi.NewErr(Fail{})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.RunTxn[sabi.Dax](base, logic)
	}
	b.StopTimer()
}

func BenchmarkDaxPrev_RunTxn_rollback_fiveDs(b *testing.B) {
	b.StopTimer()
	prev.ResetGlobals()
	defer prev.ResetGlobals()

	base := prev.NewDaxBase()
	base.SetUpLocalDaxSrc("cliargs", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("database", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("pubsub", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("json", prev_supp.FooDaxSrc{})
	base.SetUpLocalDaxSrc("env", prev_supp.FooDaxSrc{})

	type Fail struct{}

	logic := func(dax prev.Dax) prev.Err {
		conn0, err0 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "cliargs")
		conn1, err1 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "database")
		conn2, err2 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "pubsub")
		conn3, err3 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "json")
		conn4, err4 := prev.GetDaxConn[prev_supp.FooDaxConn](dax, "env")
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
		return prev.NewErr(Fail{})
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		prev.RunTxn[prev.Dax](base, logic)
	}
	b.StopTimer()
}
