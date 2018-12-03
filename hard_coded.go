package main

import (
	"fmt"
	"github.com/scorum/bitshares-go/types"
)

// Transfer a certain amount of the given asset
func ddd() {

	from := types.ObjectID{
		Space: 1,
		Type:  2,
		ID:    17,
	}

	to := types.ObjectID{
		Space: 1,
		Type:  2,
		ID:    19,
	}

	amountID := types.ObjectID{
		Space: 1,
		Type:  3,
		ID:    0,
	}

	amount := types.AssetAmount{
		AssetID: amountID,
		Amount:  10000000,
	}

	fee := types.AssetAmount{
		AssetID: amountID,
		Amount:  10000,
	}

	op := types.NewTransferOperation(from, to, amount, fee)

	fmt.Printf("%v", op)

	// fees, err := client.Database.GetRequiredFee([]types.Operation{op}, fee.AssetID.String())
	// if err != nil {
	// 	log.Println(err)
	// 	return errors.Wrap(err, "can't get fees")
	// }
	// op.Fee.Amount = fees[0].Amount

	// stx, err := client.sign([]string{key}, op)
	// if err != nil {
	// 	return err
	// }
	// return client.broadcast(stx)
}
