package main

// Minimal types so the file is self-contained.
type Request struct {
	ID int
}

func handle(req Request) {
	// pretend to process the request
}

func processRequests(ch <-chan Request) {
	// PROBLEM: This goroutine will wait forever if 'ch' is never closed,
	// causing a goroutine leak in long-running services.
	go func() {
		for req := range ch {
			handle(req)
		}
	}()
}
