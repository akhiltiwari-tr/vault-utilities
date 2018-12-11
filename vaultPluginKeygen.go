package main

import (
	"crypto/sha512"
	"fmt"
	"lib"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"github.com/vsergeev/btckeygenie/btckey"
	"golang.org/x/crypto/ripemd160"
)

// Code for creating private key using the same algorithm as used by the vault plugin.
// DerivePrivateKey Derives Private Key from mnemonic coressponding to a hard coded derivation path i.e "m/44'/69'/69'/69/69" for bitshares, to obtain private key and checks for errors.
func DerivePrivateKey(mnemonic string) (string, error) {
	// creating a private key from a mnemonic
	fmt.Println("brain key: ", mnemonic)

	// obatin private key from seed + derivation path
	// obtain seed from mnemonic and passphrase
	seed, err := lib.SeedFromMnemonic(mnemonic, "")

	btcecPrivKey, err := lib.DerivePrivateKey(seed, "m/44'/69'/69'/69/69", false)
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}

	network := &chaincfg.MainNetParams

	privateWIF, err := btcutil.NewWIF(btcecPrivKey, network, false)
	if err != nil {
		return "", err
	}

	// store private string as internal data
	PrivateKey := privateWIF.String()
	return PrivateKey, nil
}

func main() {
	brainKey := "crop fire buddy magic creek build middle digital sail state priority unhappy upper advance share"

	privKey, _ := DerivePrivateKey(brainKey)

	fmt.Println("*****", privKey)

	// converting  the private key to a btckey.PrivateKey format
	var structuredPrivKey btckey.PrivateKey
	// structuredPrivKey.FromBytes(privKey)
	structuredPrivKey.FromWIF(privKey)

	fmt.Println("#######Print results for btckey.PrivateKey format#######")
	fmt.Printf("btckey.PrivateKey.ToWIF() {BITSHARES WIF KEY}:= %v \n", structuredPrivKey.ToWIF())
	fmt.Printf("btckey.PrivateKey.ToWIFC():= %v \n", structuredPrivKey.ToWIFC())

	// retrieving public key form the private key structure.
	fmt.Println("####### Print results for btckey.PublicKey format #######")
	fmt.Printf("btckey.PrivateKeyPublicKey.ToBytes():= %v \n", structuredPrivKey.PublicKey.ToBytes())

	fmt.Println("####### Generating Bitshares' Public Key #######")
	h := ripemd160.New()
	h.Write(structuredPrivKey.PublicKey.ToBytes())
	checkSum := h.Sum(nil)
	addt := append(structuredPrivKey.PublicKey.ToBytes(), checkSum[0:4]...)
	btsPubKey := "BTS" + base58.Encode(addt)
	fmt.Println("public key", btsPubKey)

	// generating BTS address
	shaHash := sha512.New()
	shaHash.Write(structuredPrivKey.PublicKey.ToBytes())
	pubKey1 := shaHash.Sum(nil)

	mdHash := ripemd160.New()
	mdHash.Write(pubKey1)
	rawAddress := mdHash.Sum(nil)

	mdHash = ripemd160.New()
	mdHash.Write(rawAddress)
	checkSum = mdHash.Sum(nil)
	address := append(rawAddress, checkSum[0:4]...)

	result := "BTS" + base58.Encode(address)
	fmt.Println("Address: ", result)
}
