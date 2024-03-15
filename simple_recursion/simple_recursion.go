package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"syscall"
)

func getGoroutineID() int {
	buf := make([]byte, 64*1024)
	runtime.Stack(buf, true)
	goidStr := ""
	// Parse the stack trace to get the Goroutine ID
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		if strings.Contains(line, "goroutine ") {
			goidStr = line
			break
		}
	}
	var gid int
	fmt.Sscanf(goidStr, "goroutine %d ", &gid)
	return gid
}

func getThreadID() int {
	return syscall.Gettid()
}

func recursiveFunction(iteration int) {
	// Base case: if iteration reaches 10, stop recursion
	if iteration >= 2 {
		fmt.Println("Maximum iteration reached. Exiting recursion.")
		return
	}

	goid := getGoroutineID()
	tgid := getThreadID()
	fmt.Printf("TGID: %d, GOID: %d\n", tgid, goid)

	fmt.Println("Iteration:", iteration)

	// Recursive call with iteration incremented
	recursiveFunction(iteration + 1)
}

func main() {
	pid := os.Getpid()
	fmt.Println("PID:", pid)
	// Start recursion from iteration 0
	recursiveFunction(0)
}
