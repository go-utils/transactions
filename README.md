# transactions

Simply transactions

# Usage

```go
package main

import (
	"context"
	"errors"

    "github.com/go-utils/transactions"
)

var failure = errors.New("failure")

func main() {
	ctx := context.Background()

	onTransaction := transactions.OnTransaction(
		func(ctx context.Context) error {
			return failure
		},
	)

	onRollback := transactions.OnRollback(
		func(ctx context.Context, err error) error {
			if !errors.Is(err, failure) {
				return errors.New("is not `failure` error")
			}
			return nil
		},
	)

	transaction := transactions.New(onTransaction, onRollback)


	if err := transaction.Execute(ctx); err != nil {
		panic(err)
	}
}
```
