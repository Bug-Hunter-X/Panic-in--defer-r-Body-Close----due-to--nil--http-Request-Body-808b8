func handleRequest(w http.ResponseWriter, r *http.Request) {
    // ... some code ...
    defer r.Body.Close() // This line may cause error if r.Body is nil
    // ... some other code ...
}