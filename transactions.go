// Package transactions - pseudo transaction
package transactions

import (
	"context"
	"sync"
)

// OnTransaction - processing at transaction
type OnTransaction func(ctx context.Context) error

// OnRollback - processing at rollback
type OnRollback func(ctx context.Context, err error) error

// Transaction - interface
type Transaction interface {
	Execute(ctx context.Context) error
}

var _ Transaction = new(transaction)

type transaction struct {
	mtx sync.Mutex

	onTx OnTransaction
	onRb OnRollback
}

// New - constructor
func New(onTx OnTransaction, onRb OnRollback) Transaction {
	return &transaction{
		onTx: onTx,
		onRb: onRb,
	}
}

// Execute - transaction execution
func (tr *transaction) Execute(ctx context.Context) error {
	tr.mtx.Lock()
	defer tr.mtx.Unlock()

	if err := tr.onTx(ctx); err != nil {
		return tr.onRb(ctx, err)
	}

	return nil
}
