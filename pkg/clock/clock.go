package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	for i := 0; i < 50; i++ {
		measure(i)
	}
}

func measure(n int) {
	var (
		tid     = unix.Gettid()
		ctime_s unix.Timespec
		ptime_s unix.Timespec
		ttime_s unix.Timespec
		ctime_e unix.Timespec
		ptime_e unix.Timespec
		ttime_e unix.Timespec
	)

	_ = unix.ClockGettime(unix.CLOCK_MONOTONIC, &ctime_s)
	_ = unix.ClockGettime(unix.CLOCK_PROCESS_CPUTIME_ID, &ptime_s)
	_ = unix.ClockGettime(unix.CLOCK_THREAD_CPUTIME_ID, &ttime_s)

	fmt.Printf("fib: %d, ctime: %d, ptime: %d, t-%d time: %d\n", n, ctime_s.Nano(), ptime_s.Nano(), tid, ttime_s.Nano())

	fib(n)

	_ = unix.ClockGettime(unix.CLOCK_MONOTONIC, &ctime_e)
	_ = unix.ClockGettime(unix.CLOCK_PROCESS_CPUTIME_ID, &ptime_e)
	_ = unix.ClockGettime(unix.CLOCK_THREAD_CPUTIME_ID, &ttime_e)

	fmt.Printf("fib: %d, ctime-c: %d, ptime-c: %d, t-%d time-c: %d\n", n, ctime_e.Nano()-ctime_s.Nano(), ptime_e.Nano()-ptime_s.Nano(), tid, ttime_e.Nano()-ttime_s.Nano())
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
