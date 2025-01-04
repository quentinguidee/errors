# errors

A Go package designed to propagate http errors in an API codebase.

## Installation

```go
import "github.com/quentinguidee/errors"
```

## Usage

In this example, we want to create a 404 "not found" error and a 503 "service unavailable" error: 

```go
var err error
err = errors.NotFound("your resource was not found")
err = errors.ServiceUnavailable("the api cannot be accessed for the moment")
```

You can see all implemented errors in [errors.go](https://github.com/quentinguidee/errors/blob/32e33399f18a58341b2c2af0809799708d1e56e0/errors.go#L90-L214).

Then, we can propagate these errors in the application, to catch them later in an http middleware:

```go
err := c.Error()
var target *errors.HTTPError
if errors.As(err, &target) && target.Code < 500 {
    // Send the 'target' error directly as a json to the client.	
} else {
    // Log this error locally and return a generic error to hide implementation logic errors.
}
```

## LICENSE

This package is licensed under the [MIT License](./LICENSE.md).
