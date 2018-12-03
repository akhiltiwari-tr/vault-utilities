package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/vsergeev/btckeygenie/btckey"
	"golang.org/x/crypto/ripemd160"
)

func main() {

	// creating a private key from a mnemonic
	brainKey := "model margin spot canyon announce cricket lesson genre violin tuition oak love promote obscure hat glass chronic marriage spatial castle rescue sauce capable wheat"
	h := sha512.New()
	h.Write([]byte(brainKey + " " + "0"))
	pk1 := h.Sum(nil)

	hdash := sha256.New()
	hdash.Write(pk1)
	privKey := hdash.Sum(nil)
	fmt.Println("Private Key: ", privKey)

	// converting  the private key to a btckey.PrivateKey format
	var structuredPrivKey btckey.PrivateKey
	structuredPrivKey.FromBytes(privKey)

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
	fmt.Println("BTS" + base58.Encode(addt))

}
