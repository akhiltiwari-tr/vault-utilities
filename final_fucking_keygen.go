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
	brainKey := "crop fire buddy magic creek build middle digital sail state priority unhappy upper advance share"

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
	// structuredPrivKey.FromWIF("5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3")

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

// BTSFAbAx7yuxt725qSZvfwWqkdCwp9ZnUama
// BTSFAbAx7yuxt725qSZvfwWqkdCwp9ZnUama
// BTSFAbAx7yuxt725qSZvfwWqkdCwp9ZnUama
// BTSFAbAx7yuxt725qSZvfwWqkdCwp9XKEFvs
// BTSFN9r6VYzBK8EKtMewfNbfiGCr56pHDBFi

// priv key for our account pk=5Jt6AzJfm41ynUGfHV75sJvXtiRiaLuP92ra54ZmnFYiCKr9hw2
// uuid : bg780r5gouhijfg0vtq0
// address:"BTSVwfbNwjx9X3XvQKuDG6n89RCaewwWFkH",
// publicKey: "BTS89nzTzp3YGGTzMoWq398gzrFaBSpn7mBxXYQwKhHDv4PmXpbae",
