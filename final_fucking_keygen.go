package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/vsergeev/btckeygenie/btckey"
	"golang.org/x/crypto/ripemd160"
)

func genBTSWifPrivateKey(brainKey string) []byte {
	h := sha512.New()
	h.Write([]byte(brainKey + " " + "0"))
	pk1 := h.Sum(nil)

	hdash := sha256.New()
	hdash.Write(pk1)
	privKey := hdash.Sum(nil)

	return privKey
}

// GenBTSWifPrivateKey returns BTS keys from mnemonics.
func GenBTSWifPrivateKey(privKey []byte) string {
	// converting  the private key to a btckey.PrivateKey format
	var structuredPrivKey btckey.PrivateKey
	structuredPrivKey.FromBytes(privKey)
	return structuredPrivKey.ToWIF()
}

// GenBTSPublicKey returns Public key for bitshares
func GenBTSPublicKey(structuredPrivKey btckey.PrivateKey) string {
	h := ripemd160.New()
	h.Write(structuredPrivKey.PublicKey.ToBytes())
	checkSum := h.Sum(nil)
	addt := append(structuredPrivKey.PublicKey.ToBytes(), checkSum[0:4]...)
	pubKey := "BTS" + base58.Encode(addt)
	fmt.Println(pubKey)
	return "BTS" + base58.Encode(addt)
}

// Code for creating private key using the same algorithm as used by the vault plugin.

func main() {

	// creating a private key from a mnemonic
	brainKey := "hybrid easily okay write size tape clown whale dash admit leader ask weekend winner brown"

	h := sha512.New()
	h.Write([]byte(brainKey))
	pk1 := h.Sum(nil)

	hdash := sha256.New()
	hdash.Write(pk1)
	privKey := hdash.Sum(nil)
	fmt.Println("Private Key: ", privKey)

	// converting  the private key to a btckey.PrivateKey format
	var structuredPrivKey btckey.PrivateKey
	structuredPrivKey.FromBytes(privKey)
	// structuredPrivKey.FromWIF("5K44658mCyJQtmT2TtwgDJtCmduDGibq64mQ9MBBQVWtWSMLvYu")

	fmt.Println("#######Print results for btckey.PrivateKey format#######")
	fmt.Printf("btckey.PrivateKey.ToWIF() {BITSHARES WIF KEY}:= %v \n", structuredPrivKey.ToWIF())
	fmt.Printf("btckey.PrivateKey.ToWIFC():= %v \n", structuredPrivKey.ToWIFC())

	// retrieving public key form the private key structure.
	fmt.Println("####### Print results for btckey.PublicKey format #######")
	fmt.Printf("btckey.PrivateKeyPublicKey.ToBytes():= %v \n", structuredPrivKey.PublicKey.ToBytes())

	fmt.Println("####### Generating Bitshares' Public Key #######")
	h = ripemd160.New()
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

	// fmt.Println("Address: ", "BTS"+base58.Encode(address))
	result := "BTS" + base58.Encode(address)

	fmt.Println("Address: ", result)

	// var testPK ecdsa.PrivateKey

}

// hybrid easily okay write size tape clown whale dash admit leader ask weekend winner brown ==> mnemonic for nathan
