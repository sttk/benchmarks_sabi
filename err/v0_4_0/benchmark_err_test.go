package v0_4_0_test

import (
	"testing"

	go_error "github.com/sttk/benchmarks_sabi/err/go-error"
	sabi "github.com/sttk/benchmarks_sabi/err/v0_4_0"
	supp "github.com/sttk/benchmarks_sabi/err/v0_4_0/supp"
)

func BenchmarkGoError_ok(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := go_error.ReturnNilError()
		e = err
	}
	_ = e
}

func BenchmarkErr_____ok(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := supp.ReturnOkErr()
		e = err
	}
	_ = e
}

func BenchmarkGoError_ok_isOk(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnNilError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		if err == nil {
			e = err
		}
	}
	_ = e
}

func BenchmarkErr_____ok_isOk(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnOkErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		if err.IsOk() {
			e = err
		}
	}
	_ = e
}

func BenchmarkGoError_ok_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnNilError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		switch err.(type) {
		case nil:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func BenchmarkErr_____ok_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnOkErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		switch err.Reason().(type) {
		case nil:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func BenchmarkGoError_ok_ErrorString(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		str := "nil"
		s = str
	}
	_ = s
}

func BenchmarkErr_____ok_ErrorString(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnOkErr()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

func BenchmarkGoError_empty(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := go_error.ReturnEmptyError()
		e = err
	}
	_ = e
}

func BenchmarkErr_____empty(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := supp.ReturnEmptyReasonedErr()
		e = err
	}
	_ = e
}

func BenchmarkGoError_empty_isNotOk(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnEmptyError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		if err != nil {
			e = err
		}
	}
	_ = e
}

func BenchmarkErr_____empty_isNotOk(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnEmptyReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		if !err.IsOk() {
			e = err
		}
	}
	_ = e
}

func BenchmarkGoError_empty_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnEmptyError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		switch err.(type) {
		case go_error.EmptyError:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func BenchmarkErr_____empty_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnEmptyReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		switch err.Reason().(type) {
		case supp.EmptyReason:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func BenchmarkGoError_empty_ErrorString(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnEmptyError()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

func BenchmarkErr_____empty_ErrorString(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnEmptyReasonedErr()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

func BenchmarkGoError_oneField(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := go_error.ReturnOneFieldError()
		e = err
	}
	_ = e
}

func BenchmarkErr_____oneField(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := supp.ReturnOneFieldReasonedErr()
		e = err
	}
	_ = e
}

func BenchmarkGoError_oneFieldPtr(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := go_error.ReturnOneFieldErrorPtr()
		e = err
	}
	_ = e
}

func BenchmarkErr_____oneFieldPtr(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := supp.ReturnOneFieldReasonedPtrErr()
		e = err
	}
	_ = e
}

func BenchmarkGoError_oneField_isNotOk(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnOneFieldError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		if err != nil {
			e = err
		}
	}
	_ = e
}

func BenchmarkErr_____oneField_isNotOk(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnOneFieldReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		if !err.IsOk() {
			e = err
		}
	}
	_ = e
}

func BenchmarkGoError_oneField_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnOneFieldError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		switch err.(type) {
		case go_error.OneFieldError:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func BenchmarkErr_____oneField_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnOneFieldReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		switch err.Reason().(type) {
		case supp.OneFieldReason:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func BenchmarkGoError_oneField_ErrorString(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnOneFieldError()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

func BenchmarkErr_____oneField_ErrorString(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnOneFieldReasonedErr()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

func BenchmarkGoError_fiveField(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := go_error.ReturnFiveFieldError()
		e = err
	}
	_ = e
}

func BenchmarkErr_____fiveField(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := supp.ReturnFiveFieldReasonedErr()
		_ = err
	}
}

func BenchmarkGoError_fiveField_isNotOk(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnFiveFieldError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		if err != nil {
			e = err
		}
	}
	_ = e
}

func BenchmarkErr_____fiveField_isNotOk(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnFiveFieldReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		if !err.IsOk() {
			e = err
		}
	}
	_ = e
}

func BenchmarkGoError_fiveField_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnFiveFieldError()

	b.StartTimer()
	var e error
	for i := 0; i < b.N; i++ {
		switch err.(type) {
		case go_error.FiveFieldError:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func BenchmarkErr_____fiveField_typeSwitch(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnFiveFieldReasonedErr()

	b.StartTimer()
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		switch err.Reason().(type) {
		case supp.FiveFieldReason:
			e = err
		default:
			b.Fail()
		}
	}
	_ = e
}

func BenchmarkGoError_fiveField_ErrorString(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnFiveFieldError()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

func BenchmarkErr_____fiveField_ErrorString(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnFiveFieldReasonedErr()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

func BenchmarkGoError_havingCause(b *testing.B) {
	var e error
	for i := 0; i < b.N; i++ {
		err := go_error.ReturnHavingCauseError()
		e = err
	}
	_ = e
}

func BenchmarkErr_____havingCause(b *testing.B) {
	var e sabi.Err
	for i := 0; i < b.N; i++ {
		err := supp.ReturnHavingCauseReasonedErr()
		e = err
	}
	_ = e
}

func BenchmarkGoError_havingCause_ErrorString(b *testing.B) {
	b.StopTimer()
	err := go_error.ReturnHavingCauseError()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}

func BenchmarkErr_____havingCause_ErrorString(b *testing.B) {
	b.StopTimer()
	err := supp.ReturnHavingCauseReasonedErr()

	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		str := err.Error()
		s = str
	}
	_ = s
}
