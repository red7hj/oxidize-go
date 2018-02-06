package utxo_test

import (
	"reflect"
	"testing"

	"github.com/tclchiam/block_n_go/blockchain/engine/utxo"
	"github.com/tclchiam/block_n_go/identity"
)

func Test_Engine_FindUnspentOutputs(t *testing.T) {
	type args struct {
		spender *identity.Identity
	}
	tests := []struct {
		name    string
		engine  utxo.Engine
		args    args
		want    *utxo.TransactionOutputSet
		wantErr bool
	}{
		{
			name:   "",
			engine: utxo.NewAggregatorEngine(),
			args:   args{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.engine.FindUnspentOutputs(tt.args.spender)
			if (err != nil) != tt.wantErr {
				t.Errorf("Engine.FindUnspentOutputs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Engine.FindUnspentOutputs() = %v, want %v", got, tt.want)
			}
		})
	}
}
