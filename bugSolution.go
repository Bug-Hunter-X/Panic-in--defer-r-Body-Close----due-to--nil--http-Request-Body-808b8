func handleRequest(w http.ResponseWriter, r *http.Request) {
    // ... some code ...
    if r.Body != nil {
        defer r.Body.Close()
    }
    // ... some other code ...
} 