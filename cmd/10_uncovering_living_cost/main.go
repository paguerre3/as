package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/api"
)

// this is a CLI application not dockerized as the assets are too large!
/**
Use Go’s pprof for CPU and memory profiling:

// To analyze the heap usage:
go tool pprof http://localhost:6060/debug/pprof/heap

go tool pprof http://localhost:6060/debug/pprof/goroutine

// To take a CPU profile for 30 seconds:
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

Analyze the Data:

top: Displays the functions consuming the most memory.
list <function>: Shows detailed memory allocation in a specific function.
web: Generates a visual representation of memory usage in your browser.

.........................................................................
// manual snapshot
curl -o heap.prof http://localhost:6060/debug/pprof/heap

Analyze the saved file:
go tool pprof heap.prof

To view the top 10 memory consumers by allocated objects:
top10 -alloc_objects

To view memory consumption by allocated bytes:
top10 -alloc_space

//inuse_space: Memory still in use, e.g. top10 -inuse_space
//alloc_space: Total memory allocated, including memory that has been freed.

//Visualize Memory Usage
//In the interactive pprof session, run:
web
*/
// this is a CLI application not dockerized as the assets are too large!
func main() {
	enableProfiling()

	err := api.RunCommandLine()
	if err != nil {
		log.Fatal(err)
	}
}

// Enable profiling for CPU and memory.
func enableProfiling() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(swg *sync.WaitGroup) {
		defer swg.Done() // Signal readiness after starting the server
		log.Println("Starting pprof ...")

		server := &http.Server{Addr: "localhost:6060"}
		go func() {
			// Start the server
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Failed to start pprof: %v", err)
			}
		}()

		// Wait until the server is reachable
		for {
			resp, err := http.Get("http://localhost:6060/debug/pprof/")
			if err == nil && resp.StatusCode == http.StatusOK {
				resp.Body.Close()
				break // Server is ready
			}
			time.Sleep(100 * time.Millisecond) // Retry after a short delay
		}
	}(&wg)
	wg.Wait()
}
