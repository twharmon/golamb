# Golamb

[![Go Reference](https://pkg.go.dev/badge/github.com/twharmon/golamb.svg)](https://pkg.go.dev/github.com/twharmon/golamb) ![](https://github.com/twharmon/golamb/workflows/Test/badge.svg) [![](https://goreportcard.com/badge/github.com/twharmon/golamb)](https://goreportcard.com/report/github.com/twharmon/golamb) [![codecov](https://codecov.io/gh/twharmon/golamb/branch/main/graph/badge.svg?token=K0P59TPRAL)](https://codecov.io/gh/twharmon/golamb)

Golamb makes it easier to write AWS Lambda functions in Go that are invoked by API Gateway Http APIs.

## Documentation
For full documentation see [pkg.go.dev](https://pkg.go.dev/github.com/twharmon/golamb).

## Usage

### Basic
```go
package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/twharmon/golamb"
)

func handler(c golamb.Context) golamb.Responder {
	// Get a query parameter
	foo := c.Request().Query("foo")

	// Get a path parameter
	bar := c.Request().Path("bar")

	return c.Response(http.StatusOK, map[string]any{
		"foo": foo,
		"bar": bar,
	})
}

func main() {
	golamb.Start(handler)
}
```

## Contribute
Make a pull request.