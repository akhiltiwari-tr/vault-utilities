package main

import (
	"crypto/rand"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/vsergeev/btckeygenie/btckey"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	// decKey := base58.Decode("5JfGg6n1epVd53R7p1d3KRsmJAz69FUz7nwnMJaKbTW2fJ2Vx8J")
	// fmt.Println("decKey: ", decKey)
	// var pk btckey.PrivateKey
	// pk.D.SetBytes(decKey)
	// pk.Derive()

	key, _ := btckey.GenerateKey(rand.Reader)
	fmt.Println(key.ToWIF())

	h := ripemd160.New()
	h.Write(key.PublicKey.ToBytes())
	checkSum := h.Sum(nil)

	addt := append(key.PublicKey.ToBytes(), checkSum[0:4]...)
	fmt.Println("BTS" + base58.Encode(addt))
}
