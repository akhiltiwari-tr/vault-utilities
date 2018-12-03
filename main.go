package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	bitshares "github.com/scorum/bitshares-go"
	"github.com/scorum/bitshares-go/sign"
	"github.com/scorum/bitshares-go/types"
	"github.com/stretchr/testify/require"
)

const (
	testNet = "ws://0.0.0.0:11011"
)

func Transfer() {
	client, err := bitshares.NewClient(testNet)
	if err != nil {
		fmt.Printf("%v", err)
	}
	to := types.ObjectID{
		Space: 1,
		Type:  2,
		ID:    19,
	}

	from := types.ObjectID{
		Space: 1,
		Type:  2,
		ID:    17,
	}

	amountID := types.ObjectID{
		Space: 1,
		Type:  3,
		ID:    0,
	}

	amount := types.AssetAmount{
		AssetID: amountID,
		Amount:  4400000,
	}

	fee := types.AssetAmount{
		AssetID: amountID,
		Amount:  2000000,
	}

	key := "5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3"
	operations := types.NewTransferOperation(from, to, amount, fee)

	stx, err := client.Sign([]string{key}, operations)
	if err != nil {
		fmt.Println(err)
	}
	expiration := time.Now().UTC().Add(10 * time.Hour)
	// fmt.Println(expiration)
	types1 := &types.Transaction{
		RefBlockNum:    42900,
		RefBlockPrefix: 4093202820,
		Expiration:     types.Time{Time: &expiration},
	}

	stx = sign.NewSignedTransaction(types1)

	stx.PushOperation(operations)

	if err := stx.Sign([]string{key}, "3aef3997194701308d57a65214a7a015d98382ab66a9bc0d655de80842b6bfc5"); err != nil {
		fmt.Println("failed to sign the transaction")
	}

	fmt.Println("*****SIGNED TRANSACTION*****")
	spew.Dump(stx)
	fmt.Println("*****SIGNED TRANSACTION*****")

	if err = client.Broadcast(stx); err != nil {
		fmt.Println(err)
	}
}

func main() {
	// TestClient_Transfer()
	Transfer()
}

func TestClient_Transfer() {
	client, err := bitshares.NewClient(testNet)
	if err != nil {
		fmt.Printf("%v", err)
	}

	nathan, err := client.Database.LookupAccounts("nathan", 1)
	abc, err := client.Database.LookupAccounts("abc-test", 1)

	log.Printf("here ----%v--- %v\n\n", nathan, abc)

	assets, err := client.Database.LookupAssetSymbols("BTS")

	fmt.Printf("%v, %v,%v", nathan, abc, assets)

	cali4889IDActiveKey := "5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3"
	from := nathan["nathan"]
	to := abc["abc-test"]
	amount := types.AssetAmount{
		AssetID: assets[0].ID,
		Amount:  10000000,
	}
	fee := types.AssetAmount{
		AssetID: assets[0].ID,
		Amount:  10000,
	}

	fmt.Printf("%v, %v,%v, %v, %v", cali4889IDActiveKey, from, to, amount, fee)

	client.Transfer(cali4889IDActiveKey, from, to, amount, fee)
}

func TestClient_LimitOrderCreate(t *testing.T) {
	client, err := bitshares.NewClient(testNet)
	require.Nil(t, err)

	cali4889arr, err := client.Database.LookupAccounts("cali4889", 1)
	require.Nil(t, err)
	cali4889ID := cali4889arr["cali4889"]

	sellAsset, err := client.Database.LookupAssetSymbols("TEST")
	require.NoError(t, err)

	buyAsset, err := client.Database.LookupAssetSymbols("PEG.FAKEUSD")
	require.NoError(t, err)

	amSell := types.AssetAmount{
		Amount:  100,
		AssetID: sellAsset[0].ID,
	}
	minBuy := types.AssetAmount{
		Amount:  10,
		AssetID: buyAsset[0].ID,
	}

	fee := types.AssetAmount{
		AssetID: sellAsset[0].ID,
	}

	expiration := 40 * time.Hour

	cali4889IDActiveKey := "5JiTY3m9u1iPfoKsZdn18pnf26XvX2WnXFJckSiSaiUniNVzxLn"

	id, err := client.LimitOrderCreate(cali4889IDActiveKey, cali4889ID, fee, amSell, minBuy, expiration, false)
	require.NoError(t, err)

	orderID := types.MustParseObjectID(id)

	err = client.LimitOrderCancel(cali4889IDActiveKey, cali4889ID, orderID, fee)
	require.NoError(t, err)
}
