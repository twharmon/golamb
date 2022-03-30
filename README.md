# Golamb

![](https://github.com/twharmon/golamb/workflows/Test/badge.svg) [![](https://goreportcard.com/badge/github.com/twharmon/golamb)](https://goreportcard.com/report/github.com/twharmon/golamb) [![](https://gocover.io/_badge/github.com/twharmon/golamb)](https://gocover.io/github.com/twharmon/golamb)

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
	// AWS clients are lazy loaded.
	ddb := c.AWS().DynamoDB()

	output, err := ddb.GetItem(&dynamodb.GetItemInput{...})
	if err != nil {
		c.LogError("unable to get item: %s", err)
		return c.Response(http.StatusInternalServerError)
	}

	if len(output.Item) == 0 {
		c.LogWarning("item not found")
		return c.Response(http.StatusNotFound)
	}

	return c.Response(http.StatusOK, output.Item)
}

func main() {
	golamb.Start(handler)
}
```

## Contribute
Make a pull request.