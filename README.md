# StackTracer

StackTracer is a lightweight Go library that enhances error handling by automatically adding stack trace information to your errors.

## Features

- **Automatic error location**: Adds the file name and line number where the error originated
- **Simple API**: A single `Trace()` function to wrap your errors
- **No external dependencies**: Uses only the Go standard library
- **Minimal overhead**: Negligible performance impact
- **Standard error compatible**: Works with any error that implements the `error` interface

## Installation

```bash
go get github.com/RchrdHndrcks/stacktracer
```

## Basic Usage

```go
package main

import (
    "errors"
    "fmt"
    
    "github.com/RchrdHndrcks/stacktracer"
)

func main() {
    err := doSomething()
    if err != nil {
        fmt.Println(err)
        // Output: main.go:23 - something went wrong
    }
}

func doSomething() error {
    // Simulate an error
    err := errors.New("something went wrong")
    
    // Wrap the error with stack trace information
    return stacktracer.Trace(err)
}
```

## Added Value

### 1. Simplified Debugging

When working with complex applications, knowing exactly where an error originated can save hours of debugging time. StackTracer automatically adds this critical information to your error messages.

### 2. Easier Code Maintenance

By quickly identifying the exact location of errors, code maintenance becomes more efficient, especially in large codebases or when working in teams.

### 3. More Informative Logs

Error logs with stack trace information are much more useful for diagnosing issues in production, without the need to implement complex monitoring tools.

### 4. No Changes to Your Workflow

Unlike other solutions that require significant changes to how you handle errors, StackTracer integrates seamlessly with Go's standard error handling.

## Advanced Examples

### Error Propagation Through Multiple Layers

```go
func handler(w http.ResponseWriter, r *http.Request) {
    data, err := fetchData(r.URL.Query().Get("id"))
    if err != nil {
        // The error already contains stack trace information
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // ...
}

func fetchData(id string) ([]byte, error) {
    result, err := database.Query(id)
    if err != nil {
        // Add stack trace information to the error
        return nil, stacktracer.Trace(err)
    }
    return result, nil
}
```

## License

[MIT](LICENSE)
