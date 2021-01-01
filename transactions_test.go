package transactions

import (
	"context"
	"errors"
	"testing"
)

func Test_transaction_Execute(t *testing.T) {
	type fields struct {
		onTx OnTransaction
		onRb OnRollback
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success/normal",
			fields: fields{
				onTx: func(ctx context.Context) error {
					return nil
				},
				onRb: func(ctx context.Context, err error) error {
					return errors.New("failure")
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "success/rollback",
			fields: fields{
				onTx: func(ctx context.Context) error {
					return errors.New("failure")
				},
				onRb: func(ctx context.Context, err error) error {
					return nil
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "failure/rollback",
			fields: fields{
				onTx: func(ctx context.Context) error {
					return errors.New("failure")
				},
				onRb: func(ctx context.Context, err error) error {
					return errors.New("failure")
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // escape: Using the variable on range scope `tt` in loop literal
		t.Run(tt.name, func(t *testing.T) {
			transaction := New(tt.fields.onTx, tt.fields.onRb)
			if err := transaction.Execute(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
