package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

//go:noinline
func getCurrentCPUID() int {
	var cpuid int
	tid := syscall.Gettid()
	_, _, errno := syscall.Syscall(syscall.SYS_GETTID, uintptr(0), uintptr(0), uintptr(0))
	if errno == 0 {
		cpuid = int(tid) % runtime.NumCPU()
	}
	return cpuid
}

//go:noinline
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

//go:noinline
func handler(w http.ResponseWriter, r *http.Request) {
	// pid := os.Getpid()
	// cpuid := getCurrentCPUID()
	// goid := getGoroutineID()
	// tgid := getThreadID()
	fmt.Fprintf(w, "Hi there from %s!\n", r.Host)
	// fmt.Printf("HANDLER\tPID: %d, TGID: %d, GOID: %d, CPU ID: %d\n", pid, tgid, goid, cpuid)
	fmt.Printf("Handler\n")
}

//go:noinline
func Greet(name string, i, a, b int) {
	// pid := os.Getpid()
	// cpuid := getCurrentCPUID()
	// goid := getGoroutineID()
	// tgid := getThreadID()
	// fmt.Printf("GREET \tPID: %d, TGID: %d, GOID: %d, CPU ID: %d\n", pid, tgid, goid, cpuid)
	fmt.Printf("Greet\n")

	_ = add(22, 33, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

}

//go:noinline
func getThreadID() int {
	return syscall.Gettid()
}

//go:noinline
func add(i, j, k, l, m, n, o, p, q, r, s, t int) int {
	// pid := os.Getpid()
	// cpuid := getCurrentCPUID()
	// goid := getGoroutineID()
	// tgid := getThreadID()
	// fmt.Printf("ADD \tPID: %d, TGID: %d, GOID: %d, CPU ID: %d\n", pid, tgid, goid, cpuid)
	fmt.Printf("Add\n")

	k = i + j
	if k > 60 {
		return l
	}
	return m
}

func main() {
	var wg sync.WaitGroup

	// Create a channel to receive signals
	// sigCh := make(chan os.Signal, 1)
	// signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Start your server
	fmt.Println("Server starting..")
	pid := os.Getpid()
	cpuid := getCurrentCPUID()
	goid := getGoroutineID()
	tgid := getThreadID()
	fmt.Printf("MAIN \tPID: %d, TGID: %d, GOID: %d, CPU ID: %d\n", pid, tgid, goid, cpuid)

	// http.HandleFunc("/", handler)

	names := []string{"Mauros", "Lucas", "Keremee"}

	wg.Add(1)
	go func() {
		for cnt := 0; cnt < 10; cnt++ {
			Greet(names[rand.Intn(len(names))], cnt, 1, 2)
			time.Sleep(time.Second)
		}

		wg.Done()
	}()

	start := time.Now()
	_ = add(22, 33, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("add latency: %s\n", time.Since(start))

	// go func() {
	// 	log.Fatal(http.ListenAndServe(":8080", nil))
	// }()

	// Block until a signal is received
	// sig := <-sigCh
	// fmt.Println("Received signal:", sig)

	wg.Wait() // Wait until all goroutines have finished
	fmt.Println("All goroutines have finished.")
}
