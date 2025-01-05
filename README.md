# Go Http Request Body Nil Panic

This repository demonstrates a common error in Go when handling HTTP requests: a panic caused by attempting to close a `nil` `http.Request.Body`.

The `bug.go` file shows the problematic code, while `bugSolution.go` provides a robust solution.

**Problem:**
The `defer r.Body.Close()` statement is common practice but can cause a panic if `r.Body` is `nil`, which is possible in certain HTTP request scenarios.

**Solution:**
The solution involves checking for `nil` before closing the request body. This avoids the panic and gracefully handles requests with no body.

**How to run:**
1. Clone this repository.
2. Run `go run bug.go` to observe the panic. 
3. Run `go run bugSolution.go` to see the corrected and more robust code. 