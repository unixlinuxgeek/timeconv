Run test
```
$ go test -test.v >> README.md
```

Example output:
```
=== RUN   Test_SecToMin
    timeconv_test.go:14: Test Passed: 0.098876 == 0.098876
--- PASS: Test_SecToMin (0.00s)
=== RUN   Test_MinToSec
    timeconv_test.go:26: Test Passed: 242.982304 == 242.982304
--- PASS: Test_MinToSec (0.00s)
=== RUN   Test_HourToMin
    timeconv_test.go:39: Test Passed
--- PASS: Test_HourToMin (0.00s)
=== RUN   Test_HourToSec
    timeconv_test.go:50: r: 3.285551
    timeconv_test.go:52: Test Passed
--- PASS: Test_HourToSec (0.00s)
PASS
ok  	github.com/unixlinuxgeek/timeconv	0.002s
```
