package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const keySize = 2048

// public key is composed of (n, e)
type PublicKey struct {
	N *big.Int
	E int
}

// public key is composed of (n, d)
type PrivateKey struct {
	N *big.Int
	D *big.Int
}

func main() {
	// Key generation
	pub, priv := GenerateKeyPair(keySize)

	// Encrypt
	message := new(big.Int).SetBytes([]byte("Hello World! RSA!"))
	ciphertext := Encrypt(pub, message)

	// Decrypt
	decryptedMessage := Decrypt(priv, ciphertext)

	// Print
	fmt.Printf("Original Message: %s\n", message.Bytes())
	fmt.Printf("Encrypted: %x\n", ciphertext)
	fmt.Printf("Decrypted: %s\n", string(decryptedMessage.Bytes()))
}

func GenerateKeyPair(bits int) (*PublicKey, *PrivateKey) {
	eValue := int64(65537)
	E := big.NewInt(eValue)
	N := big.NewInt(0)
	D := big.NewInt(0)

	// generate the first big prime number p
	p, err := rand.Prime(rand.Reader, bits/2)
	if err != nil {
		panic(err)
	}

	// generate the second big prime number q
	q, err := rand.Prime(rand.Reader, bits/2)
	if err != nil {
		panic(err)
	}

	// calculate N = p * q
	N.Mul(p, q)

	// calculate totient = (p - 1) * (q - 1)
	pMin := new(big.Int).Sub(p, big.NewInt(1))
	qMin := new(big.Int).Sub(q, big.NewInt(1))
	totient := new(big.Int).Mul(pMin, qMin)

	// ensure that e is coprime to totient
	// we are trying to get the greatest common divisor of the E and totient
	// as long as it's not one we keep looping and increasing E by 2
	for big.NewInt(1).Cmp(new(big.Int).GCD(nil, nil, E, totient)) != 0 {
		E.Add(E, big.NewInt(2))
	}

	// calculate D = e^-1 mod totient
	D.ModInverse(E, totient)

	return &PublicKey{N, int(E.Int64())}, &PrivateKey{N, D}
}

func Encrypt(pub *PublicKey, plaintext *big.Int) *big.Int {
	// c = m^e mod n
	e := big.NewInt(int64(pub.E))
	ciphertext := new(big.Int).Exp(plaintext, e, pub.N)
	return ciphertext
}

func Decrypt(priv *PrivateKey, ciphertext *big.Int) *big.Int {
	// m = c^d mod n
	plaintext := new(big.Int).Exp(ciphertext, priv.D, priv.N)
	return plaintext
}
