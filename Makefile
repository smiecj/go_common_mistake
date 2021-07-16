test_enum:
	go test -v enum_test.go

test_benchmark:
	go test -bench=. benchmark_test.go

test_benchmark_flag:
	go test -gcflags="-m" -bench=. benchmark_test.go

test_benchmark_forbid_inline:
	go test -gcflags="-N -l -m" -bench=. benchmark_test.go

test_inline:
	go test -v inline_test.go

test_slice:
	go test -v slice_test.go